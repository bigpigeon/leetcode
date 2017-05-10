/*
120. Triangle

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

+-----------+
|[          |
|     [2],  |
|    [3,4], |
|   [6,5,7],|
|  [4,1,8,3]|
|]          |
+-----------+

The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
*/
package main

import (
	"fmt"
)

func min(numbers ...int) int {
	if len(numbers) == 0 {
		panic("function need numbers")
	}
	minimum := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < minimum {
			minimum = numbers[i]
		}
	}
	return minimum
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	cache := make([]int, len(triangle))
	cache[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		preval := cache[0]
		cache[0] = cache[0] + triangle[i][0]
		for j := 1; j < i; j++ {
			cache[j], preval = triangle[i][j]+min(cache[j], preval), cache[j]
		}
		cache[i] = triangle[i][i] + preval
		//		fmt.Println(cache)
	}
	return min(cache...)
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
	pascal_triangle := generate(10)
	fmt.Println(minimumTotal(pascal_triangle))
	custom_triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(custom_triangle))
}
