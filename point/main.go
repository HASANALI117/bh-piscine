package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printChar(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func printPoint(ptr *point) {
	printChar("x = ")
	z01.PrintRune(rune(ptr.x/10 + 48))
	z01.PrintRune(rune(ptr.x%10 + 48))
	printChar(", y = ")
	z01.PrintRune(rune(ptr.y/10 + 48))
	z01.PrintRune(rune(ptr.y%10 + 48))
	z01.PrintRune(10)
}

func main() {
	points := &point{}

	setPoint(points)
	printPoint(points)
}
