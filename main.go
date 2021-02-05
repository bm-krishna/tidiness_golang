package main

import (
	"fmt"
	"plugin"

	"github.com/bm-krishna/tidiness_golang/plugins/configuration"
)

func main() {
	path := configuration.Init()
	fmt.Println(path, "path")
	pl, err := plugin.Open(path)
	if err != nil {
		fmt.Println("pulins path not found")
	}

	symbol, err := pl.Lookup("Main")
	if err != nil {
		fmt.Println("plugin look up failed")
	}

	plfunc, ok := symbol.(func() ([]byte, error))
	if !ok {
		fmt.Println("Failed to return plugin defincation")
	}
	nameBytes, err := plfunc()
	fmt.Println(string(nameBytes), " in main")

}
