package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 9; i >= 0; i-- {
		for j := 9; j >= 0; j-- {
			if i == 0 && j == 0 {
				z01.PrintRune(rune(i) + 48)
				z01.PrintRune(rune(j) + 48)
			} else {
				z01.PrintRune(rune(i) + 48)
				z01.PrintRune(rune(j) + 48)
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
