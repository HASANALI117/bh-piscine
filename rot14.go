package piscine

func Rot14(s string) string {
	newS := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'l' {
			newS += string(s[i] + 14)
		} else if s[i] > 'l' && s[i] <= 'z' {
			newS += string(s[i] - 12)
		} else if s[i] >= 'A' && s[i] <= 'L' {
			newS += string(s[i] + 14)
		} else if s[i] > 'L' && s[i] <= 'Z' {
			newS += string(s[i] - 12)
		} else {
			newS += string(s[i])
		}
	}
	return newS
}
