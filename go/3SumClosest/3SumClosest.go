/*
16. 3Sum Closest
Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.
+----------------------------------------------------------------+
|  For example, given array S = {-1 2 1 -4}, and target = 1.     |
|                                                                |
|  The sum that is closest to the target is 2. (-1 + 2 + 1 = 2). |
+----------------------------------------------------------------+
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := 0
	min_distance := math.MaxInt64
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {
			b := nums[start]
			c := nums[end]
			sum := a + b + c
			distance := (sum - target)

			abs_distance := distance
			if abs_distance < 0 {
				abs_distance = -abs_distance
			}
			if min_distance > abs_distance {
				min_distance = abs_distance
				closest = sum
			}
			if distance == 0 {
				return closest
			} else if distance > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return closest
}

func main() {
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))
	fmt.Println(threeSumClosest([]int{1, -2, 1, 3, 2}, 1))
}
