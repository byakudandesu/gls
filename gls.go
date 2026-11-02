package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Invalid Directory")
	}
	fileList, err := os.ReadDir(currentDirectory)
	if err != nil {
		fmt.Println("Invalid Files")
	}
	for _, file := range fileList {
		if !isHidden(file.Name()) {
			fmt.Println(file.Name())
		}
	}
}

func isHidden(file string) bool {
	return strings.HasPrefix(file, ".")
}
