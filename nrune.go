package piscine

func NRune(s string, n int) rune {
	run := []rune(s)
	if n <= len(run)-1 && n >= 0 {
		return run[n]
	}
	return 0
}
