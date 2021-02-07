package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	pluginsBuilder "github.com/bm-krishna/tidiness_golang/net/http/patternrouting/plugins/builder"
	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes/builder"
)

type HandlerService struct {
	mu sync.Mutex
}
type ResourceNotFound struct {
	Message string
	Status  int
}

func PluginsRoutesConfigProvider() ([]string, map[string]string, error) {
	relativePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get relative path")
	}

	pluginsRoutesMapper, err := pluginsBuilder.PluginsConfigProvider(relativePath)
	if err != nil {
		return nil, nil, errors.New("Failde to Fetch Plugins config" + err.Error())
	}
	routesConfig, err := builder.RoutesConfigProvider(relativePath)
	if err != nil {
		return nil, nil, errors.New("Failde to Fetch Routes config" + err.Error())
	}
	return routesConfig, pluginsRoutesMapper, nil
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
	routes, pluginsRoutesMapper, err := PluginsRoutesConfigProvider()
	if err != nil {
		res.WriteHeader(http.StatusNotFound)

		res.Write([]byte(err.Error()))
		return
	}
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
