package piscine

func ToLower(s string) string {
	newS := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			newS += string(char + 32)
		} else {
			newS += string(char)
		}
	}
	return newS
}
