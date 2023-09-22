package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{-540104, -356763, -76934, 419122, 476220, 936101, 979428, 983718}

	result1 := piscine.IsSorted(piscine.IsSorted2, a1)
	result2 := piscine.IsSorted(piscine.IsSorted2, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
