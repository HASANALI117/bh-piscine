package piscine

func Index(s string, toFind string) int {
	for i := range s {
		if len(s) > len(toFind) {
			if s[i:i+len(toFind)] == toFind {
				return i
			}
		}
	}
	return -1
}
