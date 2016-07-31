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
