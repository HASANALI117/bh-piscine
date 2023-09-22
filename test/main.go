package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{-957199, -715862, -172884, -84325, -50793, 152869, 209515, 481400}

	result1 := piscine.IsSorted(piscine.IsSorted2, a1)
	result2 := piscine.IsSorted(piscine.IsSorted2, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
