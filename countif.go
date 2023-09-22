package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, char := range tab {
		if f(char) {
			count += 1
		}
	}
	return count
}
