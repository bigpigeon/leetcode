/*
84. Largest Rectangle in Histogram

Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

         _
       _| |
      | | |
      | | |  _
   _  | | |_| |
  | |_| | | | |
  |_|_|_|_|_|_|
Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
         _
       _|_|
      |/|/|
      |/|/|  _
   _  |/|/|_| |
  | |_|/|/| | |
  |_|_|/|/|_|_|
The largest rectangle is shown in the shaded area, which has area = 10 unit.
For example,
Given heights = [2,1,5,6,2,3],
return 10.

tag: stack

solution:
use a list to save grow up height and ther position
when height large than the tail with list, calculate area and remove it
*/
package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	heiList := []int{}       // save height with grow
	heiLengthList := []int{} // save line length with height
	max := 0
	for i, currHei := range heights {
		j := len(heiList) - 1
		currPos := i
		for ; j >= 0; j-- {
			// if current height larger than the tail of heiList,remove it
			if heiList[j] >= currHei {
				area := (i - heiLengthList[j]) * heiList[j]
				if area > max {
					max = area
				}
				currPos = heiLengthList[j]
			} else {
				break
			}
		}
		heiList = append(heiList[:j+1], currHei)
		heiLengthList = append(heiLengthList[:j+1], currPos)
	}
	for j := len(heiList) - 1; j >= 0; j-- {
		area := (len(heights) - heiLengthList[j]) * heiList[j]
		if area > max {
			max = area
		}
	}
	return max
}

func main() {
	//	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	//	fmt.Println(largestRectangleArea([]int{0, 1}))
	//	fmt.Println(largestRectangleArea([]int{1}))
	//	fmt.Println(largestRectangleArea([]int{2, 1, 2}))
	fmt.Println(largestRectangleArea([]int{4, 2, 3}))
	fmt.Println(largestRectangleArea([]int{5, 4, 1, 2}))
	fmt.Println(largestRectangleArea([]int{5, 4, 1, 12}))
	length := 10000
	n := make([]int, length)
	for i := 0; i < length; i++ {
		n[i] = i + 1
	}
	fmt.Println(largestRectangleArea(n))
}
