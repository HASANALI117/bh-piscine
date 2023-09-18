package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = min + i
	}
	return arr
}
