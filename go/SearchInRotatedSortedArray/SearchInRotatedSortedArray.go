/*
33. Search in Rotated Sorted Array

Suppose a sorted array is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.
*/
package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	//search max
	last, end := 0, len(nums)

	for last < end {
		mid := (end-last)/2 + last
		//		fmt.Printf("%d:%d, %d:%d %d:%d\n", last, nums[last], end, nums[end-1], mid, nums[mid])
		if nums[last] > nums[end-1] {
			if nums[last] > nums[mid] {
				end = mid
			} else {
				last = mid
			}
		} else {
			last, end = end-1, end
			break
		}
	}
	//search target
	for s, e := 0, len(nums); s < e; {
		m := (e-s)/2 + s
		absPos := (m + end) % len(nums)
		//		fmt.Printf("%d, %d, %d\n", s, e, m)
		if nums[absPos] == target {
			return absPos
		} else if nums[absPos] < target {
			s = m + 1
		} else {
			e = m
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{6, 7, 0, 1, 2, 3, 4, 5}, 3))
	fmt.Println(search([]int{6, 7, 8, 0, 1, 2, 3, 4, 5}, 25))
	fmt.Println(search([]int{6, 7, 8, 0, 1, 2, 3, 4, 5}, 4))
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7}, 25))
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7, 8}, 25))
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
}
