package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	file_path := []rune(os.Args[0])

	for _, name := range file_path[2:] {
		z01.PrintRune(name)
	}
	z01.PrintRune(10)
}
