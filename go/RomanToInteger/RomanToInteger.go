/*
13. Roman to Integer
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
*/
package main

import (
	"fmt"
)

func romanCharToInt(ch byte) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	if len(s) <= 0 {
		return 0
	}
	prev := romanCharToInt(s[0])
	curr := prev
	digit := 0
	for i := 1; i < len(s); i++ {
		curr = romanCharToInt(s[i])
		if curr <= prev {
			digit += prev
		} else {
			digit -= prev
		}
		prev = curr
	}
	digit += curr
	return digit
}

func main() {
	fmt.Println(romanToInt("DCXIX"))
	fmt.Println(romanToInt("MCMXCVI"))
	fmt.Println(romanToInt("DCXXI"))
}
