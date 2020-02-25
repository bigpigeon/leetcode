/*
680. Valid Palindrome II
Easy

Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:

Input: "aba"
Output: True

Example 2:

Input: "abca"
Output: True
Explanation: You could delete the character 'c'.

Note:

    The string will only contain lowercase characters a-z. The maximum length of the string is 50000.

*/

package main

import "fmt"

func validMustPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		r--
		l++
	}
	return true
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			if validMustPalindrome(s[l:r]) || validMustPalindrome(s[l+1:r+1]) {
				return true
			} else {
				return false
			}
		}
		r--
		l++
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abc"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abcde"))
}
