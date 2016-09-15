/*
59. Spiral Matrix II

Given an integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

For example,
Given n = 3,

You should return the following matrix:
+-------------+
|[            |
| [ 1, 2, 3 ],|
| [ 8, 9, 4 ],|
| [ 7, 6, 5 ] |
|]            |
+-------------+
*/
package main

import (
	"fmt"
)

const (
	x = 1
	y = 0
)

func generateMatrix(n int) [][]int {
	way := [][2]int{{x, 1}, {y, 1}, {x, -1}, {y, -1}}
	vertex := []int{n - 1, n - 1, 0, 1}
	solution := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		solution = append(solution, make([]int, n))
	}
	step := 0
	s := [2]int{0, 0}
	for i := 1; i <= n*n; {
		w := step % 4

		for i <= n*n {
			if s[way[w][0]] == vertex[w] {
				break
			}

			solution[s[y]][s[x]] = i
			i++
			s[way[w][0]] += way[w][1]
		}
		vertex[w] -= way[w][1]
		step++
	}
	return solution
}

func main() {
	fmt.Println(generateMatrix(0))
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(2))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
}
