/*
32. Longest Valid Parentheses

Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

For "(()", the longest valid parentheses substring is "()", which has length = 2.

Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.

*/
package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	length := 0
	lastError := -1
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				lastError = i
			} else {
				stack = stack[:len(stack)-1]
				var currLen int
				if len(stack) == 0 {
					currLen = i - lastError
				} else {
					currLen = i - stack[len(stack)-1]
				}
				if currLen > length {
					length = currLen
				}
			}
		}
	}
	return length
}

func main() {
	//	fmt.Println(longestValidParentheses("(()()"))
	//	fmt.Println(longestValidParentheses("()())"))
	//	fmt.Println(longestValidParentheses("(())()"))
	//	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("()(((()()"))
	fmt.Println(longestValidParentheses("()()()(((()()"))
	fmt.Println(longestValidParentheses("()()())(((()()"))
}
