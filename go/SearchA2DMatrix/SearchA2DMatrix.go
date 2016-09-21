/*
74. Search a 2D Matrix

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
For example,

Consider the following matrix:
+-------------------+
|[                  |
|  [1,   3,  5,  7],|
|  [10, 11, 16, 20],|
|  [23, 30, 34, 50] |
|]                  |
+-------------------+
Given target = 3, return true.
*/
package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	rowLen := len(matrix)
	colLen := len(matrix[0])
	for s, e := 0, rowLen*colLen; s < e; {
		mid := s + (e-s)/2
		currVal := matrix[mid/colLen][mid%colLen]
		if currVal == target {
			return true
		} else if currVal > target {
			e = mid
		} else {
			s = mid + 1
		}
	}
	return false
}

func main() {
	m1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(m1, 3))
	m2 := [][]int{{1, 1}}
	fmt.Println(searchMatrix(m2, 3))
}
