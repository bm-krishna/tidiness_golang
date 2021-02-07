package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Main(bytes []byte) ([]byte, error) {
	payload := make(map[string]interface{})
	fmt.Println("main in plugins")
	err := json.Unmarshal(bytes, &payload)
	if err != nil {
		return nil, errors.New("Failed to unmarshalling the payload" + err.Error())
	}
	name := ""
	nameData, found := payload["name"]
	if found {
		name = nameData.(string) + "krishna"
	}
	response := make(map[string]interface{})
	response["name"] = name
	resp, err := json.Marshal(&response)
	if err != nil {
		fmt.Println()
		return nil, errors.New("failed to Marshal resp" + err.Error())
	}
	return resp, nil
}
