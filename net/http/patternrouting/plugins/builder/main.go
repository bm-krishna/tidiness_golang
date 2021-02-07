package builder

import (
	"errors"
	"fmt"
	"io/ioutil"
	"plugin"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service string            `yaml:service`
	Paths   map[string]string `yaml:paths`
}

func PluginsBuilder(payload []byte, path string) ([]byte, error) {
	pluginPath := strings.ReplaceAll(path, ".json", "")
	pl, err := plugin.Open(pluginPath + "/main.so")
	if err != nil {
		fmt.Println("pulins path not found")
		return nil, errors.New("Plugin path not found. Failed to Build plugin" + err.Error())
	}

	symbol, err := pl.Lookup("Main")
	if err != nil {
		return nil, errors.New("Plugin look up failed." + err.Error())

	}
	// difine plugin symbol with byte argument
	pluginFunction, ok := symbol.(func([]byte) ([]byte, error))
	if !ok {
		fmt.Println("Failed to return plugin defincation")
		return nil, errors.New("Failed to return plugin defincation" + err.Error())
	}
	respBytes, err := pluginFunction(payload)
	return respBytes, nil
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
		pluginPath := relativePath + "/net/http/patternrouting/plugins/provider" + config
		routesPluginMapper[route] = pluginPath
	}
	return routesPluginMapper, nil
}
