package main

import (
	"encoding/json"
	"fmt"
)

func Main() ([]byte, error) {
	fmt.Println("main in plugins")
	details := map[string]string{
		"name": "mohan",
	}
	bytes, err := json.Marshal(&details)
	if err != nil {
		fmt.Println("failed to unmarhal", err)
		return nil, err
	}
	return bytes, nil
}
