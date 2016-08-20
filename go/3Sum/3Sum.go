/*
15. 3Sum
Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.
+----------------------------------------------------+
|For example, given array S = [-1, 0, 1, 2, -1, -4], |
|                                                    |
|A solution set is:                                  |
|[                                                   |
|  [-1, 0, 1],                                       |
|  [-1, -1, 2]                                       |
|]                                                   |
+----------------------------------------------------+
*/
package main

import (
	"fmt"
	"sort"
)

type NumsSet map[int]map[int]int

func (set NumsSet) Add(a, b, c int) {
	_, exist := set[a]
	if !exist {
		set[a] = map[int]int{}
	}
	set[a][b] = c
}

func printSelect(nums []int, a, b, c int) {
	for i, n := range nums {
		if i == a {
			fmt.Printf("(a)%d ", n)
		} else if i == b {
			fmt.Printf("(b)%d ", n)
		} else if i == c {
			fmt.Printf("(c)%d ", n)
		} else {
			fmt.Printf("%d ", n)
		}
	}
	fmt.Println("")
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	solution := NumsSet{}
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {
			b := nums[start]
			c := nums[end]
			if a+b+c == 0 {
				solution.Add(a, b, c)
				end--
			} else if a+b+c > 0 {
				end--
			} else {
				start++
			}

		}
	}
	solution_list := make([][]int, 0, len(solution))
	for k, v := range solution {
		for _k, _v := range v {
			solution_list = append(solution_list, []int{k, _k, _v})
		}
	}
	return solution_list
}

func main() {
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -5, -5, 2, 2, 3}))
	fmt.Println(threeSum([]int{-3, -3, -3, -4, -2, -2, -2, 0, 2, 2, 2, 3, 3, 3, 4}))
}
