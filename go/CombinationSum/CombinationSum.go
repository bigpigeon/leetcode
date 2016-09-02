/*
39. Combination Sum

Given a set of candidate numbers (C) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
+------------+
|[           |
| [7],       |
| [2, 2, 3]  |
|]           |
+------------+
*/
package main

import (
	"fmt"
)

func combinationSum(candidates []int, target int) [][]int {
	solution := [][]int{}
	if len(candidates) == 0 {
		return solution
	}
	v := candidates[len(candidates)-1]

	for i := 0; true; i++ {
		r := []int{}
		for d := 0; d < i; d++ {
			r = append(r, v)
		}
		if v*i == target {
			solution = append(solution, r)
		} else if v*i < target {
			ss := combinationSum(candidates[:len(candidates)-1], target-(v*i))
			for _, v := range ss {
				v = append(v, r...)
				solution = append(solution, v)
			}
		} else {
			break
		}
	}
	return solution
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{}, 7))
	fmt.Println(combinationSum([]int{1, 3, 5, 8}, 21))
}
