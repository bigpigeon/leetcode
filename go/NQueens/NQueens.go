/*
51. N-Queens

The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.



Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.

For example,
There exist two distinct solutions to the 4-queens puzzle:
+------------------------+
|[                       |
| [".Q..",  // Solution 1|
|  "...Q",               |
|  "Q...",               |
|  "..Q."],              |
|                        |
| ["..Q.",  // Solution 2|
|  "Q...",               |
|  "...Q",               |
|  ".Q.."]               |
|]                       |
+------------------------+
Merge 51-52
*/
package main

import (
	"fmt"
)

func checkSlanting(used []int, line, row int) bool {
	for uLine, uRow := range used {
		if (uLine-line) == (uRow-row) || (uLine-line) == -(uRow-row) {
			return false
		}
	}
	return true
}

func findQueensCombine(n, line int, used []int) [][]int {
	s := [][]int{}
	candidate := map[int]bool{}
	for row := 0; row < n; row++ {
		if checkSlanting(used, line, row) {
			candidate[row] = true
		}
	}
	for _, row := range used { //skip the used row
		delete(candidate, row)
	}

	if line < n-1 {
		for row, _ := range candidate {
			ret := findQueensCombine(n, line+1, append(used, row))
			s = append(s, ret...)
		}
	} else {
		for row, _ := range candidate {
			s = append(s, append(used, row))
		}
	}
	return s
}

func solveNQueens(n int) [][]string {
	s := findQueensCombine(n, 0, []int{})
	solution := [][]string{}
	lStr := ""
	for i := 0; i < n; i++ {
		lStr += "."
	}
	for _, sub := range s {
		Queens := []string{}
		for _, row := range sub {
			sCopy := []byte(lStr)
			sCopy[row] = 'Q'
			Queens = append(Queens, string(sCopy))
		}
		solution = append(solution, Queens)
	}
	return solution
}

func totalNQueens(n int) int {
	return len(solveNQueens(n))
}

func main() {
	for i := 0; i < 7; i++ {
		fmt.Println(solveNQueens(i))
		//		fmt.Println(len(solveNQueens(i)))
	}
}
