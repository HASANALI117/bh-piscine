package piscine

func IsNumeric(s string) bool {
	for _, char := range s {
		if !(rune(char) >= '0' && rune(char) <= '9') {
			return false
		}
	}
	return true
}
