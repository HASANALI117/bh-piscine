package main

import "github.com/01-edu/z01"

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"
	for _, z := range alpha {
		z01.PrintRune(z)
	}
}
