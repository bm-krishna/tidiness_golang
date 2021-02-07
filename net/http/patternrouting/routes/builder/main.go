package builder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Routes  map[string]interface{} `yaml: routes`
	Service string                 `yaml: service`
}

func RoutesConfigProvider(relativePath string) ([]string, error) {
	routers := &Config{}
	filePath := relativePath + "/net/http/patternrouting/routes/provider/index.yaml"
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("Failed to Read File path config" + err.Error())
	}
	err = yaml.Unmarshal(fileData, routers)
	if err != nil {
		return nil, errors.New("Failed to unmarshal routes data" + err.Error())
	}

	routes := routers.Routes
	var buildroutes []string
	for _, filePathConfig := range routes {
		// get json file path Routers
		buildPathByusingRex := strings.ReplaceAll(filePathConfig.(string), "$.", "")
		readRouteConfig := relativePath + "/net/http/patternrouting/routes/provider" + buildPathByusingRex
		routeFileData, err := ioutil.ReadFile(readRouteConfig)
		if err != nil {
			return nil, errors.New("Failed to read route config form provider" + err.Error())
		}
		var routeConfig map[string]string
		err = json.Unmarshal(routeFileData, &routeConfig)
		if err != nil {
			return nil, errors.New("Failed to unmarshal route data info" + err.Error())
		}
		path, found := routeConfig["path"]
		if found {
			buildroutes = append(buildroutes, path)
		}
	}
	return buildroutes, nil
}
