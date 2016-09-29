/*
85. Maximal Rectangle

Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

For example, given the following matrix:
+---------+
|1 0 1 0 0|
|1 0 1 1 1|
|1 1 1 1 1|
|1 0 0 1 0|
+---------+

Return 6.


similar 84. Largest Rectangle in Histogram
*/
package main

import (
	"fmt"
)

func printPersistTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}

//84. Largest Rectangle in Histogram  solution
func largestRectangleArea(heights []int) int {
	heiList := []int{}       // save height with grow
	heiLengthList := []int{} // save line length with height
	max := 0
	for i, currHei := range heights {
		j := len(heiList) - 1
		currPos := i
		for ; j >= 0; j-- {
			// if current height larger than the tail of heiList,remove it
			if heiList[j] >= currHei {
				area := (i - heiLengthList[j]) * heiList[j]
				if area > max {
					max = area
				}
				currPos = heiLengthList[j]
			} else {
				break
			}
		}
		heiList = append(heiList[:j+1], currHei)
		heiLengthList = append(heiLengthList[:j+1], currPos)
	}
	for j := len(heiList) - 1; j >= 0; j-- {
		area := (len(heights) - heiLengthList[j]) * heiList[j]
		if area > max {
			max = area
		}
	}
	return max
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxCol := len(matrix)
	maxRow := len(matrix[0])
	sumTable := make([][]int, maxCol)
	for i := 0; i < maxCol; i++ {
		sumTable[i] = make([]int, maxRow)
	}

	maxArea := 0

	for j := 0; j < maxRow; j++ {
		sumTable[0][j] = int(matrix[0][j] - '0')
	}
	area := largestRectangleArea(sumTable[0])
	if area > maxArea {
		maxArea = area
	}

	for i := 1; i < maxCol; i++ {
		for j := 0; j < maxRow; j++ {
			if matrix[i][j] != '0' {

				sumTable[i][j] = sumTable[i-1][j] + 1
			} else {
				sumTable[i][j] = 0
			}

		}
		area := largestRectangleArea(sumTable[i])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func main() {
	m := [][]byte{
		[]byte("10100"),
		[]byte("10111"),
		[]byte("11111"),
		[]byte("10010"),
	}
	fmt.Println(maximalRectangle(m))
	fmt.Println(maximalRectangle([][]byte{{'1'}}))
	fmt.Println(maximalRectangle([][]byte{
		[]byte("01"),
		[]byte("01"),
	}))
	fmt.Println(maximalRectangle([][]byte{
		[]byte("000"),
		[]byte("000"),
		[]byte("111"),
	}))
	fmt.Println(maximalRectangle([][]byte{
		[]byte("101001110"),
		[]byte("111000001"),
		[]byte("001100011"),
		[]byte("011001001"),
		[]byte("110110010"),
		[]byte("011111101"),
		[]byte("101110010"),
		[]byte("111010001"),
		[]byte("011110010"),
		[]byte("100111000"),
	}))
}
