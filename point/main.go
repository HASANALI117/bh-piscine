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

func (ptr *point) printChar(s string, i int) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune(rune(i/10 + 48))
	z01.PrintRune(rune(i%10 + 48))
}

func main() {
	points := &point{}
	setPoint(points)
	points.printChar("x = ", points.x)
	points.printChar(", y = ", points.y)
	z01.PrintRune(10)
}
