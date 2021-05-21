package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var s, sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	byUsingStrings()
}

func byUsingStrings() {
	var result string
	result = strings.Join(os.Args[1:], "  ")
	fmt.Println(result)
	fmt.Println(os.Args[1:])
}
