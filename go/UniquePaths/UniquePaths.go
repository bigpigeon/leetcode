/*
62. Unique Paths

A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Note: m and n will be at most 100.

solution:
2X1 = 1
2X2 = 2X1+2X1 = 2
2X3 = 2X2+2X1 = 3
2X4 = 2X2+2X2 = 4

3X2 = 2X3 = 3
3X3 = 2X3+2X3 = 6
4X3 = 2X4+3X3 = 10
4X4 = 4x3+4x3 = 20

mXn = (m-1Xn)+(n-1Xm)
...

*/
package main

import (
	"fmt"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func uniquePaths(m int, n int) int {
	m, n = minmax(m, n)
	if m <= 0 {
		return 0
	}
	if m == 1 {
		return 1
	}
	cache := map[[2]int]int{} //mXn
	for i := 2; i <= n; i++ {
		cache[[2]int{2, i}] = i
	}
	for i := 3; i <= m; i++ {
		for j := i; j <= n; j++ {
			a, b := minmax(i, j-1)
			c, d := minmax(i-1, j)
			cache[[2]int{i, j}] = cache[[2]int{a, b}] + cache[[2]int{c, d}]
			//			fmt.Println(i, j, ":", a, b, "X", c, d, "=", cache[[2]int{i, j}])
		}
	}
	return cache[[2]int{m, n}]
}

func main() {
	fmt.Println(uniquePaths(2, 3))
	fmt.Println(uniquePaths(4, 4))
	fmt.Println(1 << )
}
