/*
64. Minimum Path Sum

Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

solution like as 53. Maximum Subarray
*/
package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	//first col
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	//first row
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2},
		{3, 4},
	}))
}
