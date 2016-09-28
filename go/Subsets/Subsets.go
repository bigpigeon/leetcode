/*
78. Subsets

Given a set of distinct integers, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,3], a solution is:
+----------+
|[         |
|  [3],    |
|  [1],    |
|  [2],    |
|  [1,2,3],|
|  [1,3],  |
|  [2,3],  |
|  [1,2],  |
|  []      |
|]         |
+----------+
*/
package main

import (
	"fmt"
)

func subsetsProcess(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	solution := [][]int{}
	solution = append(solution, []int{nums[0]})
	ret := subsetsProcess(nums[1:])
	for _, v := range ret {
		solution = append(solution, append([]int{nums[0]}, v...))
	}
	solution = append(solution, ret...)
	return solution
}

func subsets(nums []int) [][]int {
	return append(subsetsProcess(nums), []int{})

}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{}))
}
