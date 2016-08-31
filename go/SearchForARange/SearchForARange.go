/*
34. Search for a Range

Given a sorted array of integers, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
*/
package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	for start, end := 0, len(nums); start < end; {
		mid := (end-start)/2 + start

		if nums[mid] == target {
			s := mid
			for ; s >= 1; s-- {
				if nums[s-1] != target {
					break
				}
			}
			e := mid
			for ; e < len(nums)-1; e++ {
				if nums[e+1] != target {
					break
				}
			}
			return []int{s, e}
		} else if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(searchRange([]int{1, 2, 3, 3, 5, 6, 9}, 3))
	fmt.Println(searchRange([]int{3, 3, 3, 3, 3, 3, 3, 3}, 3))
	fmt.Println(searchRange([]int{3, 3, 3, 3, 3, 3, 3, 3}, 25))
	fmt.Println(searchRange([]int{}, 3))
}
