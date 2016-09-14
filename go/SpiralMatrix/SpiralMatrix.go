/*
54. Spiral Matrix

Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

For example,
Given the following matrix:
+-------------+
|[            |
| [ 1, 2, 3 ],|
| [ 4, 5, 6 ],|
| [ 7, 8, 9 ] |
|]            |
+-------------+
You should return [1,2,3,6,9,8,7,4,5].
*/
package main

import (
	"fmt"
)

const (
	Xpos = 1
	Ypos = 0
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix[0]) == 1 {
		solution := []int{}
		for _, v := range matrix {
			solution = append(solution, v[0])
		}
		return solution
	}
	vertex := []int{len(matrix[0]) - 1, len(matrix) - 1, 0, 1}
	way := [][2]int{{Xpos, 1}, {Ypos, 1}, {Xpos, -1}, {Ypos, -1}}

	solution := []int{}
	start := [2]int{0, 0}
	wayStep := 0
	solution = append(solution, matrix[start[Ypos]][start[Xpos]])
	//边界问题
	for {
		w := wayStep % 4
		//		fmt.Println(wayStep, start, top[w], way[w], solution)
		if vertex[w] == start[way[w][0]] {
			break
		}
		for start[way[w][0]] != vertex[w] {
			start[way[w][0]] += way[w][1]
			solution = append(solution, matrix[start[Ypos]][start[Xpos]])
		}
		vertex[w] -= way[w][1]
		wayStep++
	}
	return solution
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(spiralOrder([][]int{
		{1},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{11, 12, 13, 14},
		{0, -1, -2, -3},
	}))
	fmt.Println(spiralOrder([][]int{
		{3},
		{2},
	}))
}
