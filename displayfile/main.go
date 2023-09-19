package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
	} else {
		filePath := os.Args[1]
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Print("Error reading file")
		}
		fmt.Print(string(fileContent))
	}
}
