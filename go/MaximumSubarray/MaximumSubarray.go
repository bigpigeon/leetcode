/*
53. Maximum Subarray

Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.



More practice:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxArray := make([]int, len(nums))
	maxArray[0] = nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		maxArray[i] = max(nums[i], nums[i]+maxArray[i-1])
		sum = max(sum, maxArray[i])
	}

	return sum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, -2, -5, -4, -1, -3}))
	fmt.Println(maxSubArray([]int{2, 5, 3, -100, 6, 4, 1, -5, 4, 8}))
}
