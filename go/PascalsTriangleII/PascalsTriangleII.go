/*
119. Pascal's Triangle II

Given an index k, return the kth row of the Pascal's triangle.

For example, given k = 3,
Return [1,3,3,1].

Note:
Could you optimize your algorithm to use only O(k) extra space?
*/
package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 2; i <= rowIndex+1; i++ {
		preval := result[0]
		for j := 1; j < i-1; j++ {
			result[j], preval = result[j]+preval, result[j]
		}
		result[i-1] = 1
	}
	return result
}

func main() {
	fmt.Println(getRow(5))
}
