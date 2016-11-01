/*
90. Subsets II

Given a collection of integers that might contain duplicates, nums, return all possible subsets.

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,2], a solution is:
+----------+
|[         |
|  [2],    |
|  [1],    |
|  [1,2,2],|
|  [2,2],  |
|  [1,2],  |
|  []      |
|]         |
+----------+

*/
package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	dup := 1
	for ; dup < len(nums) && nums[dup] == nums[0]; dup++ {

	}
	solution := subsetsWithDup(nums[dup:])
	preLen := len(solution)
	for i := 1; i <= dup; i++ {

		solution = append(solution, append([]int{}, nums[:i]...))
		for j := 1; j < preLen; j++ {
			solution = append(solution, append([]int{}, append(solution[j], nums[:i]...)...))
		}
	}
	return solution
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{1}))
	fmt.Println(subsetsWithDup([]int{1, 2}))
	fmt.Println(subsetsWithDup([]int{1, 1}))
	fmt.Println(subsetsWithDup([]int{1, 2, 3}))
	fmt.Println(subsetsWithDup([]int{9, 0, 3, 5, 7}))
}
