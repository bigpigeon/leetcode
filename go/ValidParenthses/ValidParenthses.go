/*
20. Valid Parentheses
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
*/
package main

import (
	"fmt"
)

var ParentsesPair = map[byte]byte{
	'{': '}',
	'[': ']',
	'(': ')',
}

func isValid(s string) bool {
	stack := []byte{}
	for _, c := range []byte(s) {
		switch c {
		case '{', '[', '(':
			stack = append(stack, c)
		case '}', ']', ')':
			if len(stack) > 0 {
				match, _ := ParentsesPair[stack[len(stack)-1]]
				if match != c {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{5+3(2-2)[]tt{232(232)}}"))
	fmt.Println(isValid("{5+3(2-2)[]tt{232(232)}"))
	fmt.Println(isValid("{5+3(2-2)[]tt{232(232)}}}"))
	fmt.Println(isValid("{5+3(2-2)[]tt{232(232])}}"))
	fmt.Println(isValid(""))
}
