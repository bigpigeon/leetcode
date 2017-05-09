/*
115. Distinct Subsequences

Given a string S and a string T, count the number of distinct subsequences of T in S.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).

Here is an example:
S = "rabbbit", T = "rabbit"

Return 3.

solution :
S = "rabbbit", T = "rabbit"
          _b -> i -> t
r -> a ->  b -> i -> t
         __b -> i -> t
count is 3

the rule is simple and obvious
make a list to cache character count of t in s:
traversal s
	1. if this character is same with t[0], list[0]++
	2. traversal t , if t[i+1] is same with this character, list[i+1] += list[j+1]
*/
package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	if len(t) == 0 {
		return 0
	}
	countList := make([]int, len(t))
	for i := 0; i < len(s); i++ {
		c := s[i]
		for j := len(countList) - 2; j >= 0; j-- {
			if t[j+1] == c {
				countList[j+1] += countList[j]
			}
		}
		if t[0] == c {
			countList[0]++
		}
		//		fmt.Println(countList)
	}
	return countList[len(t)-1]
}

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("rabbbrabbiit", "rabbit"))
	//	fmt.Println(numDistinct("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc"))
}
