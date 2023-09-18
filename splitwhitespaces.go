package piscine

func SplitWhiteSpaces(s string) []string {
	var currentString string
	var result []string
	for _, char := range s {
		if char == ' ' {
			result = append(result, currentString)
			currentString = ""
		} else {
			currentString += string(char)
		}
	}
	if currentString != "" {
		result = append(result, currentString)
	}
	return result
}
