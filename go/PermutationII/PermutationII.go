/*
47. Permutations II

Given a collection of numbers that might contain duplicates, return all possible unique permutations.

For example,
[1,1,2] have the following unique permutations:
+----------+
|[         |
|  [1,1,2],|
|  [1,2,1],|
|  [2,1,1] |
|]         |
+----------+
*/
package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	solution := [][]int{}
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	for _, r := range permute(nums[1:]) {
		solution = append(solution, append(r, nums[0]))
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			other := make([]int, 0, len(nums)-1)
			other = append(other, nums[:i]...)
			other = append(other, nums[i+1:]...)
			ret := permute(other)
			for _, r := range ret {
				solution = append(solution, append(r, nums[i]))
			}
		}
	}
	return solution
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return permute(nums)
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{1, 1}))
	fmt.Println(permuteUnique([]int{1, 1, 2, 2, 2}))
}
