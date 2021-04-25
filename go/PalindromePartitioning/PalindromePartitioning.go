/*
131. Palindrome Partitioning
Medium

Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.

A palindrome string is a string that reads the same backward as forward.



Example 1:

Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]

Example 2:

Input: s = "a"
Output: [["a"]]



Constraints:

    1 <= s.length <= 16
    s contains only lowercase English letters.

*/

package main

import "fmt"

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	pdList := [][]string{}
	for i := len(s) - 1; i >= 0; i-- {
		if isPalindrome(s[i:]) {
			sub := s[:i]
			if len(sub) != 0 {
				subPdList := partition(sub)
				for _, subPd := range subPdList {
					pdList = append(pdList, append(subPd, s[i:]))
				}
			} else {
				pdList = append(pdList, []string{s[i:]})
			}
		}
	}
	return pdList
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("a"))
}
