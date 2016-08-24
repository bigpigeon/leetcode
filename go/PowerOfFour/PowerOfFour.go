/*
342. Power of Four

Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example:
Given num = 16, return true. Given num = 5, return false.

Follow up: Could you solve it without loops/recursion?

Credits:
Special thanks to @yukuairoy for adding this problem and creating all test cases.
*/

package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	//equal 0b0101010101010101010101010101010101010101010101010101010101010101
	const mask int = 0x5555555555555555
	//ignore negative
	if num <= 0 {
		return false
	}
	//detect power of two
	//example: bin of eight:1000 bin of seven: 111,
	if num&(num-1) != 0 {
		return false
	}

	//detect power of four
	if (num & mask) != 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPowerOfFour(2))
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(6))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(16))
}
