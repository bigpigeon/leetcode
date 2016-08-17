/*
14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.
*/
package main

import (
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	shortest_len := MaxInt
	for _, s := range strs {
		if len(s) < shortest_len {
			shortest_len = len(s)
		}
	}
	for i := 0; i < shortest_len; i++ {
		for j := 1; j < len(strs); j++ {
			// all str compare with the first str
			if strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:shortest_len]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{"secret", "security", "securing"}))
	fmt.Println(longestCommonPrefix([]string{"abcdefg"}))
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}
