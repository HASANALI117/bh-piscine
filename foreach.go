package piscine

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for _, num := range a {
		f(num)
	}
}

func PrintNbr(i int) {
	z01.PrintRune(rune(i) + 48)
}
