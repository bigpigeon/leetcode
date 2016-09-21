/*
73. Set Matrix Zeroes

Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

Follow up:
Did you use extra space?
A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/
package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rowLen := len(matrix)
	colLen := len(matrix[0])
	rowZeroes := map[int]bool{}
	colZeroes := map[int]bool{}
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if matrix[i][j] == 0 {
				rowZeroes[i] = true
				colZeroes[j] = true
			}
		}
	}
	for k, _ := range rowZeroes {
		for j := 0; j < colLen; j++ {
			matrix[k][j] = 0
		}
	}
	for k, _ := range colZeroes {
		for i := 0; i < rowLen; i++ {
			matrix[i][k] = 0
		}
	}
}

func main() {
	a := [][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}
	setZeroes(a)
	fmt.Println(a)
}
