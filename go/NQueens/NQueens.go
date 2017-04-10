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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type UsageRange struct {
	Start, Last int
	Data        []bool
}

func InitUsageRange(start, last int) UsageRange {
	return UsageRange{start, last, make([]bool, last-start+1)}
}

func (ur *UsageRange) Usage(pos int) bool {
	return ur.Data[pos-ur.Start]
}

func (ur *UsageRange) SetUsage(pos int, val bool) {
	ur.Data[pos-ur.Start] = val
}

func formatStack(stack []int, n int) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		b := bytes.Repeat([]byte{'.'}, n)
		b[stack[i]] = 'Q'
		strs[i] = string(b)
	}
	return strs
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}
	solution := [][]string{}
	lStr := ""
	for i := 0; i < n; i++ {
		lStr += "."
	}
	xUsed := make([]bool, n)
	//sumUsed = (lim [f(x1),f(x2),...])
	//           x=0
	sumUsed := InitUsageRange(0, (n-1)*2)
	// reduceUsed = (lim [f(x1),f(x2),...])
	//               x=0
	reduceUsed := InitUsageRange(-(n - 1), (n - 1))
	stack := make([]int, n)
	var loop func([]int)
	loop = func(s []int) {
		currLine := n - len(s)
		for i := s[0]; i < n; i++ {
			for x := 1; x < len(s); x++ {
				s[x] = 0
			}
			if xUsed[i] == false &&
				sumUsed.Usage(i+currLine) == false &&
				reduceUsed.Usage(i-currLine) == false {

				xUsed[i] = true
				sumUsed.SetUsage(i+currLine, true)
				reduceUsed.SetUsage(i-currLine, true)
				//				fmt.Println(xUsed, sumUsed, reduceUsed)
				s[0] = i
				if len(s) == 1 {

					solution = append(solution, formatStack(stack, n))
				} else {
					loop(s[1:])
				}
				xUsed[i] = false
				sumUsed.SetUsage(i+currLine, false)
				reduceUsed.SetUsage(i-currLine, false)
			}
		}

	}
	loop(stack)
	return solution
}

func main() {
	for i := 1; i < 6; i++ {
		json.NewEncoder(os.Stdout).Encode(solveNQueens(i))
		fmt.Println(len(solveNQueens(i)))
	}
}
