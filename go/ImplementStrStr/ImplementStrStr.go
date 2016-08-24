/*
28. Implement strStr()

Implement strStr().

Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
*/
package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
mainRange:
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		for j := 0; j != len(needle); j++ {
			if needle[j] != haystack[i+j] {
				continue mainRange
			}
		}
		return i
	}
	return -1
}

func main() {
	fmt.Println(strStr("abcabcabcdefg", "abcd"))
	fmt.Println(strStr("a", "abcd"))
	fmt.Println(strStr("abcd", "abcd"))
}
