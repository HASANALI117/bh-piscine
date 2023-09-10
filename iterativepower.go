package piscine

func IterativePower(nb int, power int) int {
	num := 1
	for i := 1; i <= power; i++ {
		num *= nb
	}
	return num
}
