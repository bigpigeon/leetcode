/*
3. Longest Substring Without Repeating Characters

Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

import (
	"fmt"
)

func search(b []rune, target rune) int {
	for i, c := range b {
		if c == target {
			return i
		}
	}
	return -1
}

func lengthOfLongestSubstring(s string) int {
	c_list := []rune(s)
	sub := []rune{}
	max_len := 0
	for _, c := range c_list {
		index := search(sub, c)
		if index == -1 {
			sub = append(sub, c)
		} else {
			sub = append(sub[index+1:], c)
		}
		if len(sub) > max_len {
			max_len = len(sub)
		}
	}
	return max_len
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring(""))
}
