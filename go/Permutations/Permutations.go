/*
46. Permutations

Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:

+----------+
|[         |
|  [1,2,3],|
|  [1,3,2],|
|  [2,1,3],|
|  [2,3,1],|
|  [3,1,2],|
|  [3,2,1] |
|]         |
|          |
+----------+
*/
package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	solution := [][]int{}
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for i := 0; i < len(nums); i++ {
		other := make([]int, 0, len(nums)-1)
		other = append(other, nums[:i]...)
		other = append(other, nums[i+1:]...)
		ret := permute(other)
		for _, r := range ret {
			solution = append(solution, append(r, nums[i]))
		}
	}
	return solution
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
