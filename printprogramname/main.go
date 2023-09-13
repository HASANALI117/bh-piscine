package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []string(os.Args)
	for i := 0; i < len(arg); i++ {
		for _, x := range arg[i]{
			z01.PrintRune(x)
		}
	}
}