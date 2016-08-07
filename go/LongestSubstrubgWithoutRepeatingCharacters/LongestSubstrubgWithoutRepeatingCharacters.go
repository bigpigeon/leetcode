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
