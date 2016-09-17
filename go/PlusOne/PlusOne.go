/*
66. Plus One

Given a non-negative number represented as an array of digits, plus one to the number.

The digits are stored such that the most significant digit is at the head of the list.
*/
package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	carry := 1
	for j := len(digits) - 1; j >= 0; j-- {
		if carry > 0 {
			carry = (digits[j] + 1) / 10
			digits[j] = (digits[j] + 1) % 10
		} else {
			break
		}
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{}))
	fmt.Println(plusOne([]int{9, 9, 9}))
}
