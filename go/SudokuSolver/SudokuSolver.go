/*
37. Sudoku Solver

Write a program to solve a Sudoku puzzle by filling the empty cells.

Empty cells are indicated by the character '.'.

You may assume that there will be only one unique solution.

+-----------+-----+
|5 3  |  7  |     |
|6    |1 9 5|     |
|  9 8|     |  6  |
|-----+-----+-----+
|8    |  6  |    3|
|4    |8   3|    1|
|7    |  2  |    6|
|-----+-----+-----+
|  6  |     |2 8  |
|     |4 1 9|    5|
|     |  8  |  7 9|
+-----+-----+-----+

A sudoku puzzle...

+-----------+-----+
|5 3 4|6 7 8|9 1 2|
|6 7 2|1 9 5|3 4 8|
|1 9 8|3 4 2|5 6 7|
|-----+-----+-----+
|8 5 9|7 6 1|4 2 3|
|4 2 6|8 5 3|7 9 1|
|7 1 3|9 2 4|8 5 6|
|-----+-----+-----+
|9 6 1|5 3 7|2 8 4|
|2 8 7|4 1 9|6 3 5|
|3 4 5|2 8 6|1 7 9|
+-----+-----+-----+
*/
package main

import (
	"fmt"
)

func printSodoku(board [][]byte) {
	for i := 0; i < 9; i++ {
		fmt.Println(string(board[i]))
	}
}

func solveSudokuFill(i, j int, board [][]byte) bool {
	for ; i < 9; i++ {
		for ; j < 9; j++ {
			if board[i][j] == '.' {
				digits := map[byte]bool{
					'1': true,
					'2': true, '3': true,
					'4': true, '5': true,
					'6': true, '7': true,
					'8': true, '9': true,
				}
				for _i := 0; _i < 9; _i++ {
					d := board[_i][j]
					if d != '.' {
						//assume that board is valid
						delete(digits, d)
					}
				}
				for _j := 0; _j < 9; _j++ {
					d := board[i][_j]
					if d != '.' {
						//assume that board is valid
						delete(digits, d)
					}
				}
				//box
				si := (i / 3) * 3
				sj := (j / 3) * 3

				for _i := si; _i < 3+si; _i++ {
					for _j := sj; _j < 3+sj; _j++ {
						d := board[_i][_j]
						if d != '.' {
							//assume that board is valid
							delete(digits, d)
						}
					}
				}
				//match less value
				for k, _ := range digits {
					board[i][j] = k
					if solveSudokuFill(i, j, board) {
						return true
					}
				}
				board[i][j] = '.'
				return false
			}
		}
		j = 0
	}
	return true
}

func solveSudoku(board [][]byte) {
	solveSudokuFill(0, 0, board)
}

func main() {
	board := [][]byte{
		[]byte("..9748..."),
		[]byte("7........"),
		[]byte(".2.1.9..."),
		[]byte("..7...24."),
		[]byte(".64.1.59."),
		[]byte(".98...3.."),
		[]byte("...8.3.2."),
		[]byte("........6"),
		[]byte("...2759.."),
	}
	solveSudoku(board)
	printSodoku(board)
}
