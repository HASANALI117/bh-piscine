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

func printChar(s string, i int) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	ruin1 := rune(i/10 + 48)
	ruin2 := rune(i%10 + 48)
	z01.PrintRune(ruin1)
	z01.PrintRune(ruin2)
}

func main() {
	points := &point{}
	setPoint(points)
	printChar("x = ", points.x)
	printChar(", y = ", points.y)
}
