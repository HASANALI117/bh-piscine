package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) >= 0 {
			return false
		}
	}
	return true
}

func IsSorted2(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
