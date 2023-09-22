package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{-995198, -933815, -600589, -527392, 47076, 165895, 280637, 283022}

	result1 := piscine.IsSorted(piscine.IsSorted2, a1)
	result2 := piscine.IsSorted(piscine.IsSorted2, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
