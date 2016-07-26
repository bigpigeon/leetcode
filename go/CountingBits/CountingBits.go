package main

import (
	"fmt"
)

func countBits(num int) []int {
	l := make([]int, num+1)
	for i := 1; i < num+1; i++ {
		l[i] = l[i&(i-1)] + 1
	}
	return l
}

func main() {
	fmt.Println(countBits(5))
}
