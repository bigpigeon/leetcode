package datatype

func MakeRange(min, max int) []int {
	l := make([]int, 0, max-min)
	for i := min; i < max; i++ {
		l = append(l, i)
	}
	return l
}
