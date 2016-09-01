/*
36. Valid Sudoku

Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

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
A partially filled sudoku which is valid.

Note:
A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.
*/
package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	//check current board is valid
	for i := 0; i < 9; i++ {
		digits := map[byte]bool{
			'0': true, '1': true,
			'2': true, '3': true,
			'4': true, '5': true,
			'6': true, '7': true,
			'8': true, '9': true,
		}
		for j := 0; j < len(board[i]); j++ {
			d := board[i][j]
			if d != '.' {
				if digits[d] == false {
					return false
				}
				digits[d] = false
			}

		}
	}
	for j := 0; j < 9; j++ {
		digits := map[byte]bool{
			'0': true, '1': true,
			'2': true, '3': true,
			'4': true, '5': true,
			'6': true, '7': true,
			'8': true, '9': true,
		}
		for i := 0; i < 9; i++ {
			d := board[i][j]
			if d != '.' {
				if digits[d] == false {
					return false
				}
				digits[d] = false
			}

		}
	}

	for box := 0; box < 9; box++ {
		digits := map[byte]bool{
			'1': true,
			'2': true, '3': true,
			'4': true, '5': true,
			'6': true, '7': true,
			'8': true, '9': true,
		}
		si := box % 3
		sj := box / 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				d := board[i+si*3][j+sj*3]
				if d != '.' {
					if digits[d] == false {
						return false
					}
					digits[d] = false
				}
			}
		}
	}

	return true
}

func main() {
	board := [][]byte{
		[]byte(".87654321"),
		[]byte("2........"),
		[]byte("3........"),
		[]byte("4........"),
		[]byte("5........"),
		[]byte("6........"),
		[]byte("7........"),
		[]byte("8........"),
		[]byte("9........"),
	}
	fmt.Println(isValidSudoku(board))
}
