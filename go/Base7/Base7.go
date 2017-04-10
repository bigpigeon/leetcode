/*
504. Base 7

Given an integer, return a base 7 representation of it as a String.

Example 1:
Input: 100
Output: "202"
Example 2:
Input: -7
Output: "-10"
*/
package main

import (
	"fmt"
)

func convertTo7(num int) string {
	s := ""
	isNegative := false
	if num < 0 {
		num = -num
		isNegative = true
	}
	for ; num != 0; num /= 7 {
		s = fmt.Sprintf("%d", num%7) + s
	}
	if s == "" {
		return "0"
	}
	if isNegative {
		s = "-" + s
	}
	return s
}

func main() {
	fmt.Println(convertTo7(100))
	fmt.Println(convertTo7(-7))
	fmt.Println(convertTo7(0))
}
