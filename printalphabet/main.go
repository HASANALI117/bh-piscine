package main

import "github.com/01-edu/z01"

func main() {
	alpha := "abcdefghijklmnopqrstuvwxyz"

	for i := len(alpha) - 1; i > 0; i-- {
		z01.PrintRune(rune(alpha[i]))
	}
	z01.PrintRune('\n')
}
