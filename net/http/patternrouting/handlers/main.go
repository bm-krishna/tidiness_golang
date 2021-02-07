package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/plugins"
	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes"
)

type HandlerService struct {
	mu sync.Mutex
}
type ResourceNotFound struct {
	Message string
	Status  int
}

func PluginsRoutesProvider() ([]string, map[string]string) {
	relativePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get relative path")
	}
	// fetch plugins and routes config
	pluginsRoutesMapper, err := plugins.Service(relativePath)
	if err != nil {
		log.Fatal(err)
	}
	routesConfig, err := routes.Service(relativePath)
	return routesConfig, pluginsRoutesMapper
}
func HeaderBuilder(headers map[string]string, resp http.ResponseWriter) http.ResponseWriter {
	httpHeader := resp.Header()
	for key, val := range headers {
		httpHeader.Set(key, val)
	}
	return resp
}

var customHeaders = map[string]string{
	"Content-Type": "application/json",
}

// ServeHTTP is a method in Handler interface.
func (hs *HandlerService) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	hs.mu.Lock()
	defer hs.mu.Unlock()
	res = HeaderBuilder(customHeaders, res)
	routes, pluginsRoutesMapper := PluginsRoutesProvider()
	validRouter := false
	URL := strings.ReplaceAll(req.URL.Path, "/", "")
	// set Headers
	// res.Header().Set("Content-Type", "application/json")
	for _, route := range routes {
		if route == URL {
			validRouter = true
			break
		}
	}
	if !validRouter {
		res.WriteHeader(http.StatusNotFound)
		routeNotFoundJSON, err := json.Marshal(&ResourceNotFound{
			Message: "Request path not found in config",
			Status:  http.StatusNotFound,
		})
		if err != nil {
			if err != nil {
				res.Write([]byte(err.Error()))
				return
			}
		}
		res.Write(routeNotFoundJSON)
		return
	}
	// get pulgin config to handler request
	pluginConfig, found := pluginsRoutesMapper[URL]
	if !found {
		// fmt.Fprintf(res, pluginNotFound)
		res.WriteHeader(http.StatusNotFound)
		pluginNotFoundJSON, err := json.Marshal(&ResourceNotFound{
			Message: "Plugin not found to handle request",
			Status:  http.StatusNotFound,
		})
		if err != nil {
			res.Write([]byte(err.Error()))
			return
		}
		res.Write(pluginNotFoundJSON)
		return
	}
	json.NewEncoder(res).Encode(pluginConfig)
}
