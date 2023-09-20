package piscine

import (
	"github.com/01-edu/z01"
)

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}

func PrintNbr(i int) {
	z01.PrintRune(rune(i) + 48)
}
