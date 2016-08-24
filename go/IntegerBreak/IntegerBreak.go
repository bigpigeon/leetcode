/*
343. Integer Break

Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

For example, given n = 2, return 1 (2 = 1 + 1); given n = 10, return 36 (10 = 3 + 3 + 4).

Note: You may assume that n is not less than 2 and not larger than 58.

Hint:

1. There is a simple O(n) solution to this problem.
2. You may check the breaking results of n ranging from 7 to 10 to discover the regularities.
*/

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
