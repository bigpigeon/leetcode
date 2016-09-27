/*
77. Combinations

Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.

For example,
If n = 4 and k = 2, a solution is:
+--------+
|[       |
|  [2,4],|
|  [3,4],|
|  [2,3],|
|  [1,2],|
|  [1,3],|
|  [1,4],|
|]       |
+--------+
*/
package main

import (
	"fmt"
)

func combineFrom(n, m int, k int) [][]int {
	solution := [][]int{}
	if m-n+1 < k {
		return solution
	}
	if k == 1 {
		for i := n; i <= m; i++ {
			solution = append(solution, []int{i})
		}
		return solution
	}
	solution = append(solution, combineFrom(n+1, m, k)...)
	for _, v := range combineFrom(n+1, m, k-1) {
		solution = append(solution, append([]int{n}, v...))
	}
	return solution
}

func combine(n int, k int) [][]int {

	return combineFrom(1, n, k)

}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(1, 1))
	fmt.Println(combine(2, 2))
}
