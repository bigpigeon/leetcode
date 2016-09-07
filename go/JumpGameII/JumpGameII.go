/*
45. Jump Game II

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

For example:
Given array A = [2,3,1,1,4]

The minimum number of jumps to reach the last index is 2. (Jump 1 step from index 0 to 1, then 3 steps to the last index.)

Note:
You can assume that you can always reach the last index.
*/
package main

import (
	"fmt"
)

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	jumpStep := 0
	for i := 0; i < len(nums); {
		min_dsitance := len(nums) - 1 - i - nums[i]
		//		fmt.Println(i, nums[i], min_dsitance)
		if min_dsitance <= 0 {
			return jumpStep + 1
		}
		next := i
		for step := 1; step <= nums[i]; step++ {
			j := i + step
			distance := len(nums) - 1 - j - nums[j]
			//			fmt.Println(j, nums[j], distance)
			if distance < min_dsitance {
				min_dsitance = distance
				next = j
			}
		}
		jumpStep++
		if next == i {
			return -1
		}
		i = next
	}
	return jumpStep
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))
	fmt.Println(jump([]int{1, 1, 1, 1, 1, 1, 4}))
	fmt.Println(jump([]int{}))
	fmt.Println(jump([]int{7, 1, 2, 5, 5, 1}))
}
