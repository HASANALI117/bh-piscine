package main

import (
	"github.com/01-edu/z01"
)

func sort(a, b rune) bool {
	ascending := true
	if a >= b {
		ascending = false
	}
	return ascending
}

func main() {
	arr := []rune{'1', 'a', '2', 'A', '3', 'b', '4', 'C'}
	for _, i := range arr {
		sort(i, i+1)
	}
	for _, j := range arr {
		z01.PrintRune(j)
		z01.PrintRune(10)
	}
}
