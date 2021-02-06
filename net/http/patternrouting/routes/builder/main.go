package builder

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Routes  map[string]interface{} `yaml:routes`
	Service string                 `yaml:service`
}

// func main() {
// 	relativePath, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal("Failed to read relative path")
// 	}
// 	routesList := GenerateRoutePattern(relativePath)
// 	log.Println(routesList)
// }

func GenerateRoutePattern(relativePath string) []string {
	routers := &Config{}
	filePath := relativePath + "/net/http/patternrouting/routes/provider/index.yaml"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to Read file from path")
	}
	err = yaml.Unmarshal(fileData, routers)
	if err != nil {
		log.Fatal("Failed to unmarshal routes data")
	}

	routes := routers.Routes
	var buildroutes []string
	for _, filePathConfig := range routes {
		// get json file path froRegexm Routers
		buildPathByusingRex := strings.ReplaceAll(filePathConfig.(string), "$.", "")
		readRouteConfig := relativePath + "/net/http/patternrouting/routes/provider" + buildPathByusingRex
		log.Println(readRouteConfig, "readRouteConfig")
		routeFileData, err := ioutil.ReadFile(readRouteConfig)
		if err != nil {
			log.Fatal("Failed to read route config form provider")
		}
		var routeConfig map[string]string
		err = json.Unmarshal(routeFileData, &routeConfig)
		if err != nil {
			log.Fatal("Failed to unmarshal route data info")
		}
		path, found := routeConfig["path"]
		if found {
			buildroutes = append(buildroutes, path)
		}
	}
	return buildroutes
}
