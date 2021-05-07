/*
132. Palindrome Partitioning II
Hard

Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.



Example 1:

Input: s = "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

Example 2:

Input: s = "a"
Output: 0

Example 3:

Input: s = "ab"
Output: 1



Constraints:

    1 <= s.length <= 2000
    s consists of lower-case English letters only.


*/

package main

import (
	"fmt"
)

func findMinCut(v [][]bool) int {
	cache := make([]int, len(v))
	for i := 0; i < len(v); i++ {
		if v[0][i] == true {
			cache[i] = 0
		} else {
			distance := len(v)
			for j := 0; j < i; j++ {
				if v[j+1][i] && distance > cache[j]+1 {
					distance = cache[j] + 1
				}
			}
			cache[i] = distance

		}
	}
	return cache[len(cache)-1]
}

func minCut(s string) int {
	v := makeMirrorList(s)
	//printMirrorList(s, v)
	cache := make([]int, len(v)+1)
	for i := range cache {
		cache[i] = -1
	}
	return findMinCut(v)
}

func printMirrorList(s string, v [][]bool) {
	for _, v := range s {
		fmt.Printf("%c ", v)
	}
	fmt.Println()
	defer fmt.Println("----")
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v); j++ {
			if v[i][j] {
				fmt.Print(". ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

// calculate all palindrome
func makeMirrorList(_s string) [][]bool {
	v := make([][]bool, len(_s))
	for i := range v {
		v[i] = make([]bool, len(_s))
	}
	for i := len(_s) - 1; i >= 0; i-- {
		for j := i; j < len(_s); j++ {
			if _s[i] == _s[j] && (j-i < 3 || v[i+1][j-1] == true) {
				v[i][j] = true
			}
		}
	}

	return v
}

func mustEqual(a, b int) {
	if a != b {
		panic(fmt.Errorf("a %d not equal b %d", a, b))
	}
}

func main() {
	mustEqual(minCut("aab"), 1)
	mustEqual(minCut("abcbabcba"), 0)
	mustEqual(minCut("abbab"), 1)
	fmt.Println(minCut("leet"))
	fmt.Println(minCut("cbbbcc"))
	fmt.Println(minCut("aaabaa"))
	fmt.Println(minCut("ababababababababababababcbabababababababababababa"))
	fmt.Println(minCut("apjesgpsxoeiokmqmfgvjslcjukbqxpsobyhjpbgdfruqdkeiszrlmtwgfxyfostpqczidfljwfbbrflkgdvtytbgqalguewnhvvmcgxboycffopmtmhtfizxkmeftcucxpobxmelmjtuzigsxnncxpaibgpuijwhankxbplpyejxmrrjgeoevqozwdtgospohznkoyzocjlracchjqnggbfeebmuvbicbvmpuleywrpzwsihivnrwtxcukwplgtobhgxukwrdlszfaiqxwjvrgxnsveedxseeyeykarqnjrtlaliyudpacctzizcftjlunlgnfwcqqxcqikocqffsjyurzwysfjmswvhbrmshjuzsgpwyubtfbnwajuvrfhlccvfwhxfqthkcwhatktymgxostjlztwdxritygbrbibdgkezvzajizxasjnrcjwzdfvdnwwqeyumkamhzoqhnqjfzwzbixclcxqrtniznemxeahfozp"))
}
