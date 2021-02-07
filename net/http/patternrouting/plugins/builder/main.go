package builder

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service string            `yaml:service`
	Paths   map[string]string `yaml:paths`
}

func PluginsBuilder(relativePath string) {

}

func PluginsConfigProvider(relativePath string) (map[string]string, error) {
	configData, err := ioutil.ReadFile(relativePath + "/net/http/patternrouting/plugins/provider/index.yaml")
	if err != nil {
		return nil, errors.New("Faild to Read Plugins config" + err.Error())
	}
	config := &Config{}
	err = yaml.Unmarshal(configData, config)
	if err != nil {
		return nil, errors.New("Failed to unmarshal plugins config" + err.Error())
	}
	routesPluginMapper := make(map[string]string)
	for route, configPath := range config.Paths {
		config := strings.ReplaceAll(configPath, "$.", "")
		pluginPath := relativePath + "/net/http/patternrouting/handlers/provider" + config
		routesPluginMapper[route] = pluginPath
	}
	return routesPluginMapper, nil
}
