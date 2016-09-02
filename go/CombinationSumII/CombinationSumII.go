/*
40. Combination Sum II

Given a collection of candidate numbers (C) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

Each number in C may only be used once in the combination.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
A solution set is:
+-------------+
|[            |
|  [1, 7],    |
|  [1, 2, 5], |
|  [2, 6],    |
|  [1, 1, 6]  |
|]            |
+-------------+
*/
package main

import (
	"fmt"
	"sort"
)

func SolutionAppend(slice [][]int, elem []int) [][]int {
	//skip duplite element
	if len(slice) > 0 {
		last := slice[len(slice)-1]
		i := 0
		for ; i < len(last) && i < len(elem); i++ {
			if last[i] > elem[i] {
				return slice
			} else if last[i] < elem[i] {
				return append(slice, elem)
			}
		}
		if len(last) == len(elem) {
			return slice
		}
	}
	return append(slice, elem)
}

func subCombinationSum(candidates []int, target int) [][]int {
	solution := [][]int{}

	for i := 0; i < len(candidates); i++ {

		r := []int{candidates[i]}
		if candidates[i] == target {
			solution = SolutionAppend(solution, r)
		} else if candidates[i] < target {
			ss := subCombinationSum(candidates[i+1:], target-candidates[i])
			for _, s := range ss {
				s = append(r, s...)
				solution = SolutionAppend(solution, s)
			}
		} else {
			break
		}
	}
	return solution
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return subCombinationSum(candidates, target)
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 1, 2, 7, 6, 1, 5, 3}, 8))
	fmt.Println(combinationSum2([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7}, 21))
}
