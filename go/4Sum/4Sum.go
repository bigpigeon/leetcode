package main

import (
	"fmt"
	"sort"
)

type SolutionType [][]int

func (s SolutionType) Len() int {
	return len(s)
}

func (s SolutionType) Less(i, j int) bool {
	for n := 0; n != 4; n++ {
		if s[i][n] < s[j][n] {
			return true
		} else if s[i][n] > s[j][n] {
			return false
		}
	}
	return true
}

func (s SolutionType) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func compareFourSum(s1, s2 []int) bool {
	if s1 == nil && s2 == nil {
		return true
	}

	if s1 == nil || s2 == nil {
		return false
	}

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	solutions := SolutionType{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k]+nums[l] == target {
						solutions = append(solutions, []int{nums[i], nums[j], nums[k], nums[l]})
					}
				}
			}
		}
	}

	union_solutions := make([][]int, 0, len(solutions))
	if len(solutions) > 0 {
		sort.Sort(solutions)
		union_solutions = append(union_solutions, solutions[0])
		for i := 1; i < len(solutions); i++ {
			if !compareFourSum(solutions[i], solutions[i-1]) {
				union_solutions = append(union_solutions, solutions[i])
			}
		}
	}
	return union_solutions
}

func main() {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	fmt.Println(fourSum([]int{}, 0))
	fmt.Println(fourSum([]int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}, 0))
}
