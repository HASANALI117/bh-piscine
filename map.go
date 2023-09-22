package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, num := range a {
		result = append(result, f(num))
	}
	return result
}
