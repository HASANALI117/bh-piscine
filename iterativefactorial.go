package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb > 1 {
		return nb * IterativeFactorial(nb-1)
	}
	return 0
}
