package main

import (
	"fmt"
)

/*
Reverse digits of an integer.

Example1: x = 123, return 321
Example2: x = -123, return -321

click to show spoilers.

Have you thought about this?
Here are some good questions to ask before coding. Bonus points for you if you have already thought through this!

If the integer's last digit is 0, what should the output be? ie, cases such as 10, 100.

Did you notice that the reversed integer might overflow? Assume the input is a 32-bit integer, then the reverse of 1000000003 overflows. How should you handle such cases?

For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

Update (2014-11-10):
Test cases had been added to test the overflow behavior.
*/

// full-bit integer
//const MaxUint = ^uint(0)
//const MaxInt = int(MaxUint >> 1)
//const MinInt = -MaxInt - 1

//64-bit integer
//const MinInt = -1 << 63
//const MaxInt = -(MinInt + 1)

//32-bit integer
const MinInt = -1 << 31
const MaxInt = -(MinInt + 1)

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	y := 0
	for x != 0 {
		if y > MaxInt/10 || y < MinInt/10 {
			return 0
		}
		y = y*10 + x%10
		x /= 10
	}

	return y
}

func main() {
	fmt.Printf("max int:%d,min int:%d\n", MaxInt, MinInt)
	fmt.Println(reverse(-2345))
}
