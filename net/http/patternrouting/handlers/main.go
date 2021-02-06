package main

import (
	"log"
	"os"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/handlers/builder"
)

func main() {
	relaivePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to read relative path")
	}
	routesPlugnsMapperConfig, err := builder.PluginsBuilder(relaivePath)
	if err != nil {
		log.Fatal("Failed to build PluginsBuilder")
	}
	log.Println(routesPlugnsMapperConfig)
}

// package handlers

// import (
// 	"net/http"
// 	"regexp"
// )

// type HandlerProvider struct {
// 	routes []string
// }

// func (handlerProvider *HandlerProvider) handlerBuilder(paths []string, handler http.Handler) {
// 	// handlerProvider.routes
// }

// type route struct {
// 	pattern *regexp.Regexp
// 	handler http.Handler
// }

// type RegexpHandler struct {
// 	routes []*route
// }

// func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler) {
// 	h.routes = append(h.routes, &route{pattern, handler})
// }

// func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request)) {
// 	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
// }

// func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	for _, route := range h.routes {
// 		if route.pattern.MatchString(r.URL.Path) {
// 			route.handler.ServeHTTP(w, r)
// 			return
// 		}
// 	}
// 	// no pattern matched; send 404 response
// 	http.NotFound(w, r)
// }
