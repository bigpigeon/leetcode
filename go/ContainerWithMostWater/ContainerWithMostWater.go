/*
Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container.

example:
vartical list =>
[1, 2, 3, 3, 0, 0, 0, 2] conver to coordinate:
^
|
|     | |
|   | | |       |
| | | | |       |
+----------------->
  0 1 2 3 4 5 6 7

Container most water coordinate pair (1, x1), (7, x7)
The most container water is (7-1)*2 = 12
*/
package main

import (
	"fmt"
)

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	most_water := 0
	for left < right {
		min_height := height[left]
		width := right - left
		if height[left] > height[right] {
			min_height = height[right]
			right--
		} else {
			left++
		}
		container := width * min_height
		if most_water < container {
			most_water = container
		}
	}
	return most_water
}

func main() {
	fmt.Println(maxArea([]int{1, 2, 3, 3, 0, 0, 0, 1}))
	fmt.Println(maxArea([]int{1, 2, 3, 3, 0, 0, 0, 2}))
}
