package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// init values
	showHidden := false
	directory := "."

	// check if there are arguments
	if len(os.Args) > 1 {
		arguments := os.Args[1:]
		for _, argument := range arguments {
			if isFlag(argument) {
				switch argument {
				case "-a":
					showHidden = true
				default:
					fmt.Println("Unknown flag: ", argument)
					return
				}
			} else {
				directory = argument
			}
		}
	}
	// retrive all files under directroy
	fileList, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Invalid Files")
		os.Exit(1)
	}

	for _, file := range fileList {
		// show all files when showHidden
		if showHidden {
			fmt.Println(file.Name())
			continue
		}
		// show non hidden when not showHidden
		if !isHidden(file.Name()) {
			fmt.Println(file.Name())
		}
	}
}

// Checks if the file passed in is a dotfile
func isHidden(file string) bool {
	return strings.HasPrefix(file, ".")
}

// checks if argument is a flag
func isFlag(arg string) bool {
	return strings.HasPrefix(arg, "-")
}

// implement -l
