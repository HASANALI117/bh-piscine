package piscine

func StrLen(s string) int {
	var str []rune
	for i := 0; i < len(s); i++ {
		str = append(str, rune(s[i]))
	}
	return len(str)
}
