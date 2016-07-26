package main

import (
	"fmt"
)

func integerBreak(n int) int {
	switch n {
	case 1, 2:
		return 1
	case 3:
		return 2
	default:
		nums := n / 3
		other := 0
		sum := 1
		switch n % 3 {
		case 2:
			other = 2
		case 1:
			other = 4
			nums--
		case 0:
			other = 1
		}
		for ; nums > 0; nums-- {
			sum *= 3
		}
		return sum * other

	}

}

func main() {
	for i := 1; i != 100; i++ {
		fmt.Println(integerBreak(i))
	}

}
