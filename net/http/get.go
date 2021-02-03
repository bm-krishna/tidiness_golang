package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.mocki.io/v1/b043df5a")
	if err != nil {
		fmt.Println(err, "error")
	}
	respData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err, "err in ioutil.ReadAll")
		os.Exit(1)
	}
	// fmt.Printf("%s", respData)
	var respMap interface{}
	err = json.Unmarshal(respData, &respMap)
	if err != nil {
		fmt.Println(err, "err in unmarshal")
		os.Exit(1)
	}
	fmt.Println(respMap)
}
