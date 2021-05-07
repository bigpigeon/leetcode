/*
1745. Palindrome Partitioning IV
Hard

Given a string s, return true if it is possible to split the string s into three non-empty palindromic substrings. Otherwise, return false.​​​​​

A string is said to be palindrome if it the same string when reversed.



Example 1:

Input: s = "abcbdd"
Output: true
Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.

Example 2:

Input: s = "bcbddxy"
Output: false
Explanation: s cannot be split into 3 palindromes.



Constraints:

    3 <= s.length <= 2000
    s​​​​​​ consists only of lowercase English letters.


*/

package main

import "fmt"

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

type Area struct {
	Start, End int
}

func (a Area) Len() int {
	return a.End - a.Start + 1
}

func sameWord(s string) bool {
	if len(s) < 2 {
		return true
	}
	w := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != w {
			return false
		}
	}
	return true
}

func checkPartitioning(s string) bool {
	v := makeMirrorList(s)
	cache := make([][]Area, len(v)-1)
	for i := 0; i < len(v)-1; i++ {
		if v[0][i] == true {
			cache[i] = []Area{{0, i}}
		} else {
			var dc []Area
			for j := 0; j <= i; j++ {
				dc = append(dc, Area{j, j})
			}
			for j := 0; j < i; j++ {
				if v[j+1][i] && len(dc) > len(cache[j])+1 {
					dc = make([]Area, 0, len(cache[j])+1)
					dc = append(dc, Area{j + 1, i})
					dc = append(dc, cache[j]...)
				}
			}
			cache[i] = dc
		}
	}
	// find p = 3
	last := len(v) - 1
	for j := 0; j < last; j++ {
		if v[j+1][last] {
			if len(cache[j])+1 == 3 {
				return true
			} else if len(cache[j])+1 == 2 {
				// one cut is not zero even number
				if (last+1-(j+1) == 2) || (last+1-(j+1) > 2 && sameWord(s[j+1:last+1])) {
					return true
				}

				if cache[j][0].Len() == 2 || (cache[j][0].Len() > 2 && sameWord(s[cache[j][0].Start:cache[j][0].End+1])) {
					return true
				}

			} else if len(cache[j])+1 == 1 {
				if (last+1-(j+1) > 2) || cache[j][0].Len() > 2 {
					return true
				}
			}

		}
	}

	return false
}

func main() {
	//fmt.Println(checkPartitioning("abcbdd"))
	fmt.Println(checkPartitioning("acab"))
	fmt.Println(checkPartitioning("acaa"))
	fmt.Println(checkPartitioning("acbc"))
	fmt.Println(checkPartitioning("aaax"))
}
