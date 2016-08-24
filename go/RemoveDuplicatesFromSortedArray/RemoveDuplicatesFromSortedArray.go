/*
26. Remove Duplicates from Sorted Array

Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

*/
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	nums_len := len(nums)
	curr_len := nums_len
	for i := 1; i < nums_len; i++ {
		if nums[i] == nums[i-1] {
			tag := i
			curr_len--
			for i += 1; i < nums_len; i++ {
				// find next not duplicate value
				if nums[i] != nums[tag] {
					// move not duplicate values position to tag
					for j := 0; j+i < nums_len; j++ {
						nums[i+j], nums[tag+j] = nums[tag+j], nums[i+j]
					}
					// reset current index
					i = tag
					break
				}
				curr_len--
			}
			// reset range end point
			nums_len = curr_len
		}
	}
	return nums_len
}

func main() {
	n1 := []int{}
	fmt.Println(removeDuplicates(n1), n1)
	n2 := []int{1, 1, 2}
	fmt.Println(removeDuplicates(n2), n2)
	n3 := []int{1, 2, 3}
	fmt.Println(removeDuplicates(n3), n3)
	n4 := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(n4), n4)
	n5 := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(n5), n5)
	n6 := []int{1, 1}
	fmt.Println(removeDuplicates(n6), n6)
}
