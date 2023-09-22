package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{-983690, -808955, -487942, -364693, 244344, 273682, 588124, 780714}

	result1 := piscine.IsSorted(piscine.IsSorted2, a1)
	result2 := piscine.IsSorted(piscine.IsSorted2, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
