package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		filePath := os.Args[1]
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file")
		}
		fmt.Println(string(fileContent))
	}
}
