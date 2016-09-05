/*
44. Wildcard Matching

Implement wildcard pattern matching with support for '?' and '*'.

+----------------------------------------------------------------------+
|'?' Matches any single character.                                     |
|'*' Matches any sequence of characters (including the empty sequence).|
|                                                                      |
|The matching should cover the entire input string (not partial).      |
|                                                                      |
|The function prototype should be:                                     |
|bool isMatch(const char *s, const char *p)                            |
|                                                                      |
|Some examples:                                                        |
|isMatch("aa","a") → false                                             |
|isMatch("aa","aa") → true                                             |
|isMatch("aaa","aa") → false                                           |
|isMatch("aa", "*") → true                                             |
|isMatch("aa", "a*") → true                                            |
|isMatch("ab", "?*") → true                                            |
|isMatch("aab", "c*a*b") → false                                       |
+----------------------------------------------------------------------+
*/
package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	i := 0
	j := 0
	last_i := len(s)
	last_j := len(p)
	for i < len(s) {
		//		fmt.Println(s[i:], p[j:])
		if j < len(p) && p[j] == '*' {
			j++
			if j == len(p) {
				return true
			}
			last_i = i
			last_j = j
		} else {
			if j < len(p) && (p[j] == '?' || s[i] == p[j]) {
				i++
				j++
			} else if last_i < len(s) {

				last_i++
				i = last_i
				j = last_j
			} else {
				return false
			}
		}
	}
	for ; j < len(p); j++ {
		if p[j] != '*' {
			break
		}
	}
	return j == len(p)
}

func main() {
	fmt.Println(isMatch("abcd", "**b"))
	fmt.Println(isMatch("abcd", "**d"))
	fmt.Println(isMatch("abcd", "a?cd"))
	fmt.Println(isMatch("abcd", "*"))
	fmt.Println(isMatch("aasssswcb", "a*z*b"))
	fmt.Println(isMatch("aasssswcb", "a*w*b"))
	fmt.Println(isMatch("abefcdgiescdfimde", "ab*cd?i*de"))
	fmt.Println(isMatch("", "*"))
	fmt.Println(isMatch("ttaaxxbdabcdxxbcaabdabcdbcd", "**aa*bdabcd*bcd"))
	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb",
		"**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

}
