package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

// return p1 is bigger than p2?
func Compare(p1 Point, p2 Point) bool {
	subtract := p1.X - p2.X
	if subtract == 0 {
		return p1.Y-p2.Y > 0
	} else {
		return subtract > 0
	}
}

func SearchRoad(board [][]byte, road []Point, word string) bool {
	if len(word) == 0 || len(road) == 0 {
		return false
	}
	cent := road[len(road)-1]
	if board[cent.X][cent.Y] == word[0] {
		if len(word) == 1 {
			return true
		}

		// all might select road
		run_select := []Point{
			{cent.X, cent.Y + 1},
			{cent.X, cent.Y - 1},
			{cent.X - 1, cent.Y},
			{cent.X + 1, cent.Y},
		}

	mainRange:
		for _, r := range run_select {
			//the run point can't out of the board range
			if r.X >= len(board) || r.Y >= len(board[0]) || r.X < 0 || r.Y < 0 {
				continue
			}
			//not in exist path
			for _, e := range road {
				if e == r {
					continue mainRange
				}
			}

			result := SearchRoad(board, append(road, r), word[1:])
			if result {
				return true
			}

		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {
			road := []Point{{i, j}}
			result := SearchRoad(board, road, word)
			if result {
				return result
			}
		}

	}
	return false
}

func main() {
	fmt.Println(
		exist([][]byte{
			[]byte("ABCE"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		}, "ABCCED",
		))
	fmt.Println(
		exist([][]byte{
			[]byte("ABCE"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		}, "SEE",
		))
	fmt.Println(
		exist([][]byte{
			[]byte("ABCE"),
			[]byte("SFCS"),
			[]byte("ADEE"),
		}, "ABCB",
		))
	fmt.Println(
		exist([][]byte{
			[]byte("a"),
		}, "a",
		))
}
