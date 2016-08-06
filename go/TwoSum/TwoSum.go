package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 1, 5}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
