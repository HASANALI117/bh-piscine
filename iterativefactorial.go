package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb > 1 {
		num := 1
		for i := nb; i > 1; i-- {
			num *= i
		}
		return num
	}
	return 0
}
