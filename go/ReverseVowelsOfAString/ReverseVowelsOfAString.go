package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	byte_s := []byte(s)
	start, end := 0, len(byte_s)
RootRange:
	for {
	startRange:
		for {
			if start >= end {
				break RootRange
			}
			switch byte_s[start] {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				break startRange
			}
			start++
		}
	endRange:
		for {
			if start >= end {
				break RootRange
			}
			switch byte_s[end-1] {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				break endRange
			}
			end--
		}
		byte_s[start], byte_s[end-1] = byte_s[end-1], byte_s[start]
		start++
		end--
	}
	return string(byte_s)
}

func main() {
	fmt.Printf("%s\n", reverseVowels("hello"))
	fmt.Printf("%s\n", reverseVowels("leetcode"))
}
