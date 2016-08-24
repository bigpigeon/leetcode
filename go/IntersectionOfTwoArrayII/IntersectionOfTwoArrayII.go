/*
350. Intersection of Two Arrays II

Given two arrays, write a function to compute their intersection.

Example:
Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].

Note:
Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:
1. What if the given array is already sorted? How would you optimize your algorithm?
2. What if nums1's size is small compared to nums2's size? Which algorithm is better?
3. What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/
package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i_1 := 0
	i_2 := 0

	intersection := []int{}
	for i_1 < len(nums1) && i_2 < len(nums2) {
		if nums1[i_1] > nums2[i_2] {
			i_2++
		} else if nums2[i_2] > nums1[i_1] {
			i_1++
		} else {
			intersection = append(intersection, nums1[i_1])
			i_1++
			i_2++
		}
	}
	return intersection
}

func main() {

	rst := intersect([]int{3, 2, 1, 3}, []int{3, 4, 5, 3})
	fmt.Println(rst)
}
