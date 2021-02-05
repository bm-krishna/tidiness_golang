package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Hits int64 `yaml:"hits"`
	Time int64 `yaml:"time"`
}

func (c *conf) getConf() *conf {
	CurrentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CurrentDir)
	yamlFile, err := ioutil.ReadFile(CurrentDir + "/core/yaml/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()
	fmt.Println(c)
}
