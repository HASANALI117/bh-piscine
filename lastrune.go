package piscine

func LastRune(s string) rune {
	run := []rune(s)
	return run[len(run)-1]
}