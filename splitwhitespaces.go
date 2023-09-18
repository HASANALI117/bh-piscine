package piscine

func SplitWhiteSpaces(s string) []string {
	var currentString string
	var result []string
	for _, char := range s {
		if char == ' ' || char == 9 || char == 10 {
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
