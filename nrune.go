package piscine

func NRune(s string, n int) rune {
	run := []rune(s)
	if n <= len(run) && n >= 0 {
		return run[n-1]
	}
	return 0
}
