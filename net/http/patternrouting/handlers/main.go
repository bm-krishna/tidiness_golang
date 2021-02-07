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
		log.Println("Relative paths")
	}
	pluginsMappers := plugins.Service(relativePath)
	routes := routes.Service(relativePath)
	log.Println(pluginsMappers, routes)
}
