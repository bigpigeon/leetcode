/*
48. Rotate Image

You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Follow up:
Could you do this in-place?
*/
package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func rotate(matrix [][]int) {
	if len(matrix) <= 1 {
		return
	}
	length := len(matrix)
	for l := length; l >= 0; l -= 2 {
		for i := 0; i < l-1; i++ {
			excursion := (length - l) / 2
			_lu := Point{excursion, i + excursion}
			_ru := Point{i + excursion, l - 1 + excursion}
			_rd := Point{l - 1 + excursion, l - i - 1 + excursion}
			_ld := Point{l - i - 1 + excursion, excursion}
			//			fmt.Println(_lu, _ru, _rd, _ld, length-l)
			//			fmt.Println(matrix[_lu.X][_lu.Y], matrix[_ru.X][_ru.Y], matrix[_rd.X][_rd.Y], matrix[_ld.X][_ld.Y])
			matrix[_lu.X][_lu.Y],
				matrix[_ru.X][_ru.Y],
				matrix[_rd.X][_rd.Y],
				matrix[_ld.X][_ld.Y] =
				matrix[_ld.X][_ld.Y],
				matrix[_lu.X][_lu.Y],
				matrix[_ru.X][_ru.Y],
				matrix[_rd.X][_rd.Y]
			//			fmt.Println(matrix[_lu.X][_lu.Y], matrix[_ru.X][_ru.Y], matrix[_rd.X][_rd.Y], matrix[_ld.X][_ld.Y])
		}
	}

}

func main() {

	mCollection := [][][]int{
		[][]int{},
		[][]int{
			[]int{11, 12, 13, 14},
			[]int{21, 22, 23, 24},
			[]int{31, 32, 33, 34},
			[]int{41, 42, 43, 44},
		},
		[][]int{
			[]int{11, 12, 13, 14, 15},
			[]int{21, 22, 23, 24, 25},
			[]int{31, 32, 33, 34, 35},
			[]int{41, 42, 43, 44, 45},
			[]int{51, 52, 53, 54, 55},
		},
	}
	for _, m := range mCollection {
		rotate(m)
		fmt.Println(m)
	}

}
