package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	max := ""
	for i := 0; i != len(s); i++ {
		j := 0
		for ; i-j > -1 && i+j < len(s); j++ {
			if s[i+j] != s[i-j] {
				break
			}
		}
		if 2*j+1 > len(max) {
			max = s[i-j+1 : i+j]
		}

	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			j := 0
			for ; i-j > -1 && i+1+j < len(s); j++ {
				if s[i+1+j] != s[i-j] {
					break
				}
			}
			if 2*j+1 > len(max) {
				max = s[i-j+1 : i+j+1]
			}
		}
	}
	return max
}

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("xabcbswe"))
	fmt.Println(longestPalindrome("yurqwabcbaqabcdefedcba"))
	fmt.Println(longestPalindrome("xuwaebbeaccopi"))
	fmt.Println(longestPalindrome("xuwabccccbaopi"))
}
