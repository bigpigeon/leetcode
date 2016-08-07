package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1)+len(nums2) == 0 {
		return 0.0
	}
	numsMerge := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0
	if i != len(nums1) && j != len(nums2) {

		for {
			if nums1[i] <= nums2[j] {
				numsMerge = append(numsMerge, nums1[i])
				i++
				if i >= len(nums1) {
					break
				}
			} else {
				numsMerge = append(numsMerge, nums2[j])
				j++
				if j >= len(nums2) {
					break
				}
			}
		}
	}
	for ; i < len(nums1); i++ {
		numsMerge = append(numsMerge, nums1[i])

	}
	for ; j < len(nums2); j++ {
		numsMerge = append(numsMerge, nums2[j])
	}
	fmt.Println(numsMerge)

	if len(numsMerge)%2 == 0 {
		return float64(numsMerge[len(numsMerge)/2]+numsMerge[len(numsMerge)/2-1]) / 2.0
	}
	return float64(numsMerge[len(numsMerge)/2])

}

func main() {
	fmt.Println(findMedianSortedArrays(
		[]int{1, 3, 5, 7, 9},
		[]int{2, 4, 6, 8, 10},
	))
	fmt.Println(findMedianSortedArrays(
		[]int{1, 3},
		[]int{2},
	))
}
