/*
81. Search in Rotated Sorted Array II

Follow up for "Search in Rotated Sorted Array":
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?

Write a function to determine if a given target is in the array.
*/
package main

import (
	"fmt"
)

func search(nums []int, target int) bool {
	//find last number with sorted array
	last, end := 0, len(nums)
	for last < end {
		mid := (end-last)/2 + last
		if nums[last] > nums[end-1] {
			if nums[last] > nums[mid] {
				end = mid
			} else {
				last = mid
			}
		} else if nums[last] < nums[end-1] {
			last, end = end-1, end
			break
		} else {
			last++
		}
	}
	fmt.Println(last)

	//find target
	for s, e := 0, len(nums); s < e; {
		m := (e-s)/2 + s
		absPos := (m + end) % len(nums)
		if nums[absPos] == target {
			return true
		} else if nums[absPos] < target {
			s = m + 1
		} else {
			e = m
		}
	}
	return false
}

func main() {
	//	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 2))
	//	fmt.Println(search([]int{4, 5, 6, 6, 7, 0, 1, 2}, 8))
	fmt.Println(search([]int{1, 3, 1, 1, 1}, 3))
}
