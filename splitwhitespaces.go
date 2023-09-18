package piscine

func SplitWhiteSpaces(s string) []string {
	var currentStr string
	var result []string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == 10 || s[i] == 9 {
			if currentStr != "" {
				result = append(result, currentStr)
				currentStr = ""
			}
		} else {
			currentStr += string(s[i])
		}
	}

	if currentStr != "" {
		result = append(result, currentStr)
	}
	return result
}
