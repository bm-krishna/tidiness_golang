package builder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Routers struct {
	Routes  map[string]interface{} `yaml:routes`
	Service string                 `yaml:service`
}

// func main() {
// 	routesList := GenerateRoutePattern()
// 	log.Println(routesList)
// }

func GenerateRoutePattern() []string {
	// return root path name with corresponding current directory
	CurrentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// pwd := os.Getenv("PWD")
	// log.Println("pwd")
	// log.Println(pwd)
	routers := &Routers{}
	filePath := CurrentDir + "/provider/index.yaml"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to Read file from path")
	}
	err = yaml.Unmarshal(fileData, routers)
	if err != nil {
		log.Fatal("Failed to unmarshal routes data")
	}

	routes := routers.Routes
	service := routers.Service
	var buildroutes []string
	for _, filePathConfig := range routes {
		// get json file path froRegexm Routers
		buildPathByusingRex := strings.ReplaceAll(filePathConfig.(string), "$.", "")
		readRouteConfig := CurrentDir + "/" + service + buildPathByusingRex
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
