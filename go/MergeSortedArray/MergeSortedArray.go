/*
88. Merge Sorted Array

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.
*/
package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	curr := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[curr] = nums1[i]
			i--
		} else {
			nums1[curr] = nums2[j]
			j--
		}
		curr--
	}
	for j >= 0 {
		nums1[curr] = nums2[j]
		curr--
		j--
	}
}

func main() {
	n1 := []int{1, 3, 5, 7, 9, 0, 0, 0, 0, 0}
	n2 := []int{2, 4, 6, 8, 10}
	merge(n1, 5, n2, 5)
	fmt.Println(n1)
}
