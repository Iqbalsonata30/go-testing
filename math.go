package gotesting

func Absolute(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}