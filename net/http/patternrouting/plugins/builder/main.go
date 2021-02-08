package builder

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"plugin"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Service string            `yaml:service`
	Paths   map[string]string `yaml:paths`
}

func PluginsCompiler(path string) error {
	path = strings.ReplaceAll(path, ".json", "")
	destinationPath := path + "/main.so"
	pluginPath := path + "/main.go"
	fmt.Println(destinationPath)
	fmt.Println(pluginPath)
	command := "go build -buildmode=plugin -o " + destinationPath + " " + pluginPath
	cmd := exec.Command(command)
	result, err := cmd.Output()
	if err != nil {
		fmt.Println("in plugins Compiler")
		return errors.New("Failed to Compile Plugins" + err.Error())
	}
	output := string(result[:])
	log.Println(output)
	return nil
}
func PluginsBuilder(payload []byte, path string) ([]byte, error) {
	// compile code
	err := PluginsCompiler(path)
	if err != nil {
		return nil, errors.New(err.Error())
	}
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
