/*
41. First Missing Positive

Given an unsorted integer array, find the first missing positive integer.

For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.

Your algorithm should run in O(n) time and uses constant space.
*/
package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	nums_len := len(nums)
	missingList := make([]bool, nums_len)
	for _, v := range nums {
		if v < nums_len+1 && v > 0 {
			missingList[v-1] = true
		}
	}
	i := 0
	for ; i < nums_len; i++ {
		if missingList[i] == false {
			break
		}
	}
	return i + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{}))
}
