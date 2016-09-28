/*
80. Remove Duplicates from Sorted Array II

Follow up for "Remove Duplicates":
What if duplicates are allowed at most twice?

For example,
Given sorted array nums = [1,1,1,2,2,3],

Your function should return length = 5, with the first five elements of nums being 1, 1, 2, 2 and 3. It doesn't matter what you leave beyond the new length.
*/
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	duplicates := false
	currNum := nums[0]
	move := 0
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		nums[i-move], nums[i] = nums[i], nums[i-move]
		if v == currNum {
			if duplicates {
				move++
			} else {
				duplicates = true
			}
		} else {
			currNum = v
			duplicates = false
		}
	}
	return len(nums) - move
}

func main() {
	a := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(a))
	fmt.Println(a)
	b := []int{1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 5, 5, 5, 6, 8, 9}
	fmt.Println(removeDuplicates(b))
	fmt.Println(b)

	fmt.Println(removeDuplicates([]int{}))
}
