/*
118. Pascal's Triangle

Given numRows, generate the first numRows of Pascal's triangle.

For example, given numRows = 5,
Return
+------------+
|[           |
|     [1],   |
|    [1,1],  |
|   [1,2,1], |
|  [1,3,3,1],|
| [1,4,6,4,1]|
|]           |
+------------+
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

func generate(numRows int) [][]int {
	solution := [][]int{}
	for i := 0; i < min(numRows, 2); i++ {
		line := make([]int, i+1)
		line[0], line[i] = 1, 1
		solution = append(solution, line)
	}
	for i := 2; i < numRows; i++ {
		line := make([]int, i+1)
		line[0], line[i] = 1, 1
		preline := solution[i-1]
		for j := 1; j < i; j++ {
			line[j] = preline[j] + preline[j-1]
		}
		solution = append(solution, line)
	}
	return solution
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(generate(i))
	}
}
