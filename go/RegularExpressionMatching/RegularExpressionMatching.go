/*
10. Regular Expression Matching

Implement regular expression matching with support for '.' and '*'.
+-----------------------------------------------------------------+
|'.' Matches any single character.                                |
|'*' Matches zero or more of the preceding element.               |
|                                                                 |
|The matching should cover the entire input string (not partial). |
|                                                                 |
|The function prototype should be:                                |
|bool isMatch(const char *s, const char *p)                       |
|                                                                 |
|Some examples:                                                   |
|isMatch("aa","a") → false                                        |
|isMatch("aa","aa") → true                                        |
|isMatch("aaa","aa") → false                                      |
|isMatch("aa", "a*") → true                                       |
|isMatch("aa", ".*") → true                                       |
|isMatch("ab", ".*") → true                                       |
|isMatch("aab", "c*a*b") → true                                   |
+-----------------------------------------------------------------+
*/

package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	m := p[len(p)-1]
	if m == '*' {
		m = p[len(p)-2]
		if m == '.' {
			for i := len(s); i > -1; i-- {
				if isMatch(s[:i], p[:len(p)-2]) == true {
					return true
				}
			}
		} else {
			for i := len(s); i > -1; i-- {
				if isMatch(s[:i], p[:len(p)-2]) == true {
					return true
				} else if i > 0 && s[i-1] != m {
					return false
				}
			}
		}
		return false
	} else {
		if len(s) == 0 {
			return false
		}
		if m == '.' {
			return isMatch(s[:len(s)-1], p[:len(p)-1])
		} else {
			if s[len(s)-1] == m {
				return isMatch(s[:len(s)-1], p[:len(p)-1])
			}
			return false
		}
	}
}

func main() {

	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", ".*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println("custom test")
	fmt.Println(isMatch("aaa", "aaaa"))
	fmt.Println(isMatch("a", ".*"))
	fmt.Println(isMatch("", ".*"))
	fmt.Println(isMatch("aaa", "ab*a"))
}
