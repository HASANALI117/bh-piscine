package piscine

func StrRev(s string) string {
	var revS string
	for i := len(s) - 1; i >= 0; i-- {
		revS += string(s[i])
	}
	return revS
}
