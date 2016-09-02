/*
42. Trapping Rain Water

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

For example,
Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.

 ^
 |
 |
 |
 |             |
 |     | X X X | | X |
 | | X | | X | | | | | |
 +------------------------>

he above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
*/
package main

import (
	"fmt"
)

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	capacity := 0
	maxRight := height[0]
	maxLeft := height[len(height)-1]
	for i, j := 0, len(height)-1; i < j-1; {
		if maxRight < maxLeft {
			i++
			if height[i] < maxRight {
				capacity += maxRight - height[i]
			} else {
				maxRight = height[i]
			}
		} else {
			j--
			if height[j] < maxLeft {
				capacity += maxLeft - height[j]
			} else {
				maxLeft = height[j]
			}
		}
	}
	return capacity
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 0, 2, 1, 2, 1}))
}
