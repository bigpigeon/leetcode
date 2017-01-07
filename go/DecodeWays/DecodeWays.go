/*
91. Decode Ways

A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26

Given an encoded message containing digits, determine the total number of ways to decode it.

For example,
Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).

The number of ways decoding "12" is 2.


*/

package main

import (
	"fmt"
	"strconv"
)

func Atoi(s string) int {
	d, _ := strconv.Atoi(s)
	return d
}

func numDecodings(s string) int {
	//	fmt.Print(s, " ")
	if len(s) == 0 {
		return 0
	}
	allowZero := false
	pre := 0
	curr := 1
	for i := 0; i < len(s); i++ {
		//		fmt.Printf("digit%c\n", s[i])
		if s[i] == '0' {
			if allowZero == false || Atoi(s[i-1:i+1]) > 26 {
				return 0
			}
			curr, pre = pre, 0
			allowZero = false
		} else if i > 0 && Atoi(s[i-1:i+1]) > 26 {
			curr, pre = curr, curr
			allowZero = true
		} else {
			curr, pre = curr+pre, curr
			allowZero = true
		}
		//		fmt.Println(curr, pre, i)
	}
	return curr
}

func PrintAndDecoding(s string) {
	fmt.Println("str:", s, " ", numDecodings(s))
}

func main() {
	PrintAndDecoding("1")
	PrintAndDecoding("0")
	PrintAndDecoding("01")
	PrintAndDecoding("12")
	PrintAndDecoding("27")
	PrintAndDecoding("101")
	PrintAndDecoding("123")
	PrintAndDecoding("1231")
	PrintAndDecoding("12312")
	PrintAndDecoding("123120")
	PrintAndDecoding("123111231")
}
