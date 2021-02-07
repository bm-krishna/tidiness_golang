package main

import (
	"log"
	"os"

	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/plugins"
	"github.com/bm-krishna/tidiness_golang/net/http/patternrouting/routes"
)

func main() {
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
	log.Println(pluginsRoutesMapper, routesConfig)

}
