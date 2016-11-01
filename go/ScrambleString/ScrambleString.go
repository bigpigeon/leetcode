/*
87. Scramble String

Given a string s1, we may represent it as a binary tree by partitioning it to two non-empty substrings recursively.

Below is one possible representation of s1 = "great":

    great
   /    \
  gr    eat
 / \    /  \
g   r  e   at
           / \
          a   t

To scramble the string, we may choose any non-leaf node and swap its two children.

For example, if we choose the node "gr" and swap its two children, it produces a scrambled string "rgeat".

    rgeat
   /    \
  rg    eat
 / \    /  \
r   g  e   at
           / \
          a   t

We say that "rgeat" is a scrambled string of "great".

Similarly, if we continue to swap the children of nodes "eat" and "at", it produces a scrambled string "rgtae".

    rgtae
   /    \
  rg    tae
 / \    /  \
r   g  ta  e
       / \
      t   a

We say that "rgtae" is a scrambled string of "great".

Given two strings s1 and s2 of the same length, determine if s2 is a scrambled string of s1.
*/
package main

import (
	"fmt"
)

//dp solution
/*dp definition
* dp[k][i][j]
*    a) s1 start form i
*    b) s2 start form j
*    c) k is length of substring
 */
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	length := len(s1)
	dp := map[int]map[int]map[int]bool{}
	//initialization k=1
	dp[1] = map[int]map[int]bool{}
	for i := 0; i < length; i++ {
		dp[1][i] = map[int]bool{}
		for j := 0; j < length; j++ {
			dp[1][i][j] = (s1[i] == s2[j])
		}
	}
	for k := 2; k <= length; k++ {
		dp[k] = map[int]map[int]bool{}
		for i := 0; i < length-k+1; i++ {
			dp[k][i] = map[int]bool{}
			for j := 0; j < length-k+1; j++ {
				dp[k][i][j] = false
				for divk := 1; divk < k && dp[k][i][j] == false; divk++ {
					dp[k][i][j] = (dp[divk][i][j] && dp[k-divk][i+divk][j+divk]) ||
						(dp[divk][i][j+k-divk] && dp[k-divk][i+divk][j])
				}
			}
		}
	}
	return dp[length][0][0]
}

func main() {
	fmt.Println(isScramble("", ""))
	fmt.Println(isScramble("AB", "BA"))
	fmt.Println(isScramble("great", "rgtae"))
	fmt.Println(isScramble("great", "rteag"))
	fmt.Println(isScramble("AB", "AA"))
	fmt.Println(isScramble("abab", "aabb"))
}
