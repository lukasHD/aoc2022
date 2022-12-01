package lib

func AddArray(numbs ...int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}
