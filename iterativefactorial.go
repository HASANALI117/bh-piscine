package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 {
		return 1
	} else if nb > 1 {
		num := 1
		for i := nb; i > 1; i-- {
			num *= i
		}
		return num
	} else {
		return 0
	}
}
