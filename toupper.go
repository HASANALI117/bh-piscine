package piscine

func ToUpper(s string) string {
	newS := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			newS += string(char - 32)
		} else {
			newS += string(char)
		}
	}
	return newS
}
