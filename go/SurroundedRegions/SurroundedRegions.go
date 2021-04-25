/*
 130. Surrounded Regions
Medium

Given an m x n matrix board containing 'X' and 'O', capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.



Example 1:

Input: board = [
["X","X","X","X"],
["X","O","O","X"],
["X","X","O","X"],
["X","O","X","X"]
]
Output: [
["X","X","X","X"],
["X","X","X","X"],
["X","X","X","X"],
["X","O","X","X"]
]
Explanation: Surrounded regions should not be on the border,
which means that any 'O' on the border of the board are not flipped to 'X'.
Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
Two cells are connected if they are adjacent cells connected horizontally or vertically.

Example 2:

Input: board = [["X"]]
Output: [["X"]]



Constraints:

    m == board.length
    n == board[i].length
    1 <= m, n <= 200
    board[i][j] is 'X' or 'O'.


*/

package main

import (
	"fmt"
)

// group only storage edit with edge, example with below graph, the first line and second line last 'g' in one group
// the second line middle 'g' independent as group
// the third line twice 'g' independent as group
// their relationship g1()->g2->g3
// .....ggggggggg
// .......g.....g
// ......gg......
// ..............
// ..............

type AStar struct {
	stack  []int
	width  int
	height int
}

func (c *AStar) Put(i, j int) {
	c.stack = append(c.stack, i*c.width+j)
}

func (c *AStar) Pop() (int, int) {
	if len(c.stack) == 0 {
		return -1, -1

	}
	val := c.stack[len(c.stack)-1]
	c.stack = c.stack[:len(c.stack)-1]
	return val / c.width, val % c.width
}

func (c *AStar) Iter(f func(i, j int) bool) {
	for len(c.stack) != 0 {
		c.Step(f)
	}
}

func (c *AStar) Step(f func(i, j int) bool) {
	x, y := c.Pop()

	// iter up
	{
		i, j := x-1, y
		if i >= 1 && f(i, j) {
			c.Put(i, j)
		}
	}
	//down
	{
		i, j := x+1, y
		if i < c.height-1 && f(i, j) {
			c.Put(i, j)
		}
	}
	// left
	{
		i, j := x, y-1
		if j >= 1 && f(i, j) {
			c.Put(i, j)
		}
	}
	// right
	{
		i, j := x, y+1
		if j < c.width-1 && f(i, j) {
			c.Put(i, j)
		}
	}
}

type PointSet struct {
	group  map[int]bool
	width  int
	height int
}

func (c *PointSet) Set(i, j int) {
	c.group[i*c.width+j] = true
}

func (c *PointSet) Get(i, j int) bool {
	return c.group[i*c.width+j]
}

// pos:
//  top(t)
//left(l) right(r)
//  bottom(b)
func iterBoardEdge(board [][]byte, x, y, w, h int, f func(i, j int, pos byte)) {
	if len(board) == 0 {
		return
	}

	for i := x; i < h; i++ {
		f(i, 0, 'l')
		f(i, w-1, 'r')
	}
	for j := y; j < w; j++ {
		f(0, j, 't')
		f(h-1, j, 'b')
	}
}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m := len(board)
	n := len(board[0])
	xSet := &PointSet{width: n, height: m, group: map[int]bool{}}

	iterBoardEdge(board, 0, 0, n, m, func(i, j int, pos byte) {
		if board[i][j] == 'O' {
			xSet.Set(i, j)
			astar := AStar{
				stack:  nil,
				width:  n,
				height: m,
			}
			astar.Put(i, j)
			astar.Iter(func(i, j int) bool {
				if xSet.Get(i, j) == false && board[i][j] == 'O' {
					xSet.Set(i, j)
					return true
				}
				return false
			})
		}
	})
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if board[i][j] == 'O' && xSet.Get(i, j) == false {
				board[i][j] = 'X'
			}
		}
	}
}

func printBoard(b [][]byte) {

	for i := range b {
		for j := range b[i] {
			fmt.Printf("%c ,", b[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	{
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		solve(board)
		printBoard(board)
	}
	{
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		solve(board)
		printBoard(board)
	}
	{
		board := [][]byte{
			{'X', 'X', 'X', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'X', 'O', 'X'},
			{'X', 'O', 'O', 'X'},
			{'X', 'O', 'X', 'X'},
			{'X', 'O', 'X', 'X'},
		}
		solve(board)
		printBoard(board)
	}
}
