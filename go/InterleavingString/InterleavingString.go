/*
97. Interleaving String

Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

For example,
Given:
s1 = "aabcc",
s2 = "dbbca",

When s3 = "aadbbcbcac", return true.
When s3 = "aadbbbaccc", return false.

solution:
make a grid(n X m) to cache s3 path
n = len(s1)+1
m = len(s2)+1
grid[i][j] = (grid[i][j-1] == true && s3[i + j - 1] == s2[j-1]) ||
             (grid[i-1][j] == true && s3[i + j - 1] == s1[i-1])

result was grid last item value

e.g
s1 = "aabcc" s2 = "dbbca" s3 = "aadbbcbcac"
+---------------+
|0 0 d b b c a  |
|0   x x x x x  |
|a * x x x x x  |
|a * * * * * x  |
|b x * * x * x  |
|c x x * * * *  |
|c x x x * x *  |
+---------------+
*/
package main

import (
	"fmt"
)

func printGrid(grid [][]bool) {

	for _, v := range grid {
		fmt.Println(v)
	}
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if len(s3) == 0 {
		return true
	}
	grid := make([][]bool, len(s1)+1)
	for i, _ := range grid {
		grid[i] = make([]bool, len(s2)+1)
	}
	for i := 0; i < len(s1) && s3[i] == s1[i]; i++ {
		grid[i+1][0] = true
	}

	for j := 0; j < len(s2) && s3[j] == s2[j]; j++ {
		grid[0][j+1] = true
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			s3curr := i + j - 1
			grid[i][j] = (grid[i][j-1] == true && s3[s3curr] == s2[j-1]) ||
				(grid[i-1][j] == true && s3[s3curr] == s1[i-1])

		}
	}
	printGrid(grid)
	return grid[len(s1)][len(s2)]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("", "", ""))
}
