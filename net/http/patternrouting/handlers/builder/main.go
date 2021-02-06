package builder

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service string            `yaml:service`
	Paths   map[string]string `yaml:paths`
}

func PluginsBuilder(relativePath string) (map[string]string, error) {
	configData, err := ioutil.ReadFile(relativePath + "/net/http/patternrouting/handlers/provider/index.yaml")
	if err != nil {
		log.Fatal("Faild to Read Plugins config")
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configData, config)
	if err != nil {
		log.Fatal("Failed to unmarshal plugins config")
		return nil, err
	}
	routesPluginMapper := make(map[string]string)
	for route, configPath := range config.Paths {
		config := strings.ReplaceAll(configPath, "$.", "")
		log.Println(route, config, "config")
		pluginPath := relativePath + "/net/http/patternrouting/handlers/provider" + config
		routesPluginMapper[route] = pluginPath
	}
	return routesPluginMapper, nil
}

// func main() {
// 	relativePath, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal("Failed to read relative path")
// 	}
// 	PluginsBuilder(relativePath)
// }
