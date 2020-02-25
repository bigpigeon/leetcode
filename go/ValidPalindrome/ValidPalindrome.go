/*
125. Valid Palindrome

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:

Input: "race a car"
Output: false
*/

package main

import (
	"fmt"
)

func toLower(c byte) byte {
	if c <= 'Z' && c >= 'A' {
		return 'a' + (c - 'A')
	}
	return c
}

func isAlphaNum(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l <= r {
		if isAlphaNum(s[l]) == false {
			l++
			continue
		}
		if isAlphaNum(s[r]) == false {
			r--
			continue
		}
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		r--
		l++
	}
	return true
}

func main() {
	//fmt.Println(isPalindrome(""))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(isPalindrome("race a car"))
}
