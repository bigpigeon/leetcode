/*
76. Minimum Window Substring

Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

For example,
S = "ADOBECODEBANC"
T = "ABC"
Minimum window is "BANC".

Note:
If there is no such window in S that covers all characters in T, return the empty string "".

If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.
reference:
https://github.com/haoel/leetcode/blob/master/algorithms/cpp/minimumWindowSubstring/minimumWindowSubstring.cpp
*/
package main

import (
	"fmt"
)

func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}
	tHash := map[byte]int{}
	for i := 0; i < len(t); i++ {
		tHash[t[i]]++
	}
	sHash := map[byte]int{}
	letterFound := 0
	begin := 0
	start := 0
	last := len(s) - 1
	for i := 0; i < len(s); i++ {
		if _, ok := tHash[s[i]]; ok {
			sHash[s[i]]++
			if sHash[s[i]] <= tHash[s[i]] {
				letterFound++
			}
			if letterFound >= len(t) {
				for tHash[s[begin]] == 0 || sHash[s[begin]] > tHash[s[begin]] {
					if sHash[s[begin]] > tHash[s[begin]] {
						sHash[s[begin]]--
					}
					begin++
				}
				if last-start > i-begin {
					start = begin
					last = i
				}
			}
		}

	}
	if letterFound < len(t) {
		return ""
	}
	return s[start : last+1]
}

func main() {
	//	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	//	fmt.Println(minWindow("", "ABC"))
	//	fmt.Println(minWindow("ADOBECODEBANC", ""))
	//	fmt.Println(minWindow("adobecodebanc", "abcda"))
	fmt.Println(minWindow("ask_not_what_your_country_can_do_for_you_ask_what_you_can_do_for_your_country", "ask_country"))
	fmt.Println(minWindow("How_are_you_balabala_areyou_balabala", "areyou"))
}
