package piscine

func Any(f func(string) bool, a []string) bool {
	for _, char := range a {
		if f(char) {
			return true
		}
	}
	return false
}
