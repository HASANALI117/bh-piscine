package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	}
	if nb > 1 && nb < 40 {
		num := 1
		for i := 1; i <= nb; i++ {
			num *= i
		}
		return num
	}
	return 0
}
