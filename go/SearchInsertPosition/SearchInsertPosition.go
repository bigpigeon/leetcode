/*
35. Search Insert Position

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Here are few examples.
[1,3,5,6], 5 → 2
[1,3,5,6], 2 → 1
[1,3,5,6], 7 → 4
[1,3,5,6], 0 → 0
*/
package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)
	for start < end {
		mid := (end-start)/2 + start
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
}
