/*
22. Generate Parentheses
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:
+-----------+
|[          |
|  "((()))",|
|  "(()())",|
|  "(())()",|
|  "()(())",|
|  "()()()" |
|]          |
+-----------+

*/
package main

import (
	"fmt"
)

func generate(left, right int, s string) []string {
	if (right <= 0) && (left <= 0) {
		return []string{s}
	}
	l := []string{}

	if left > 0 {
		l = append(l, generate(left-1, right, s+"(")...)
	}
	if (right > 0) && (right > left) {
		l = append(l, generate(left, right-1, s+")")...)
	}
	return l
}

func generateParenthesis(n int) []string {
	return generate(n, n, "")
}

func main() {
	fmt.Println(generateParenthesis(3))
}
