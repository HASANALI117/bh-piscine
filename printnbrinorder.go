package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(rune(0) + 48)
		return
	}
	if n < 0 {
		return
	}
	var intArray []int
	for n > 0 {
		intArray = append(intArray, n%10)
		n /= 10
	}
	for i := 0; i < len(intArray)-1; i++ {
		if intArray[i] > intArray[i+1] {
			intArray[i], intArray[i+1] = intArray[i+1], intArray[i]
		}
	}

	for _, digit := range intArray {
		z01.PrintRune(rune(digit) + 48)
	}
}
