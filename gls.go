package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileList, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Invalid Files")
		return
	}

	for _, file := range fileList {

		if len(os.Args) > 1 {
			argument := os.Args[1]
			if showAll(argument) {
				fmt.Println(file.Name())
				continue
			}
		}
		if !isHidden(file.Name()) {
			fmt.Println(file.Name())
		}
	}
}

// Checks if the file passed in is a dotfile
func isHidden(file string) bool {
	return strings.HasPrefix(file, ".")
}

func showAll(arg string) bool {
	return arg == "-a"
}
