/*
63. Unique Paths II

Follow up for "Unique Paths":

Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.

For example,
There is one obstacle in the middle of a 3x3 grid as illustrated below.
+----------+
|[         |
|  [0,0,0],|
|  [0,1,0],|
|  [0,0,0] |
|]         |
+----------+
The total number of unique paths is 2.

Note: m and n will be at most 100.
*/
package main

import (
	"fmt"
)

func uniquePathsWithObstaclesStep(obstacleGrid [][]int, step [2]int) int {
	x, y := step[0], step[1]
	sum := 0
	flag := 0
	if obstacleGrid[y][x] == 1 {
		return 0
	}
	if x < len(obstacleGrid[0])-1 {
		right := obstacleGrid[y][x+1]
		if right < 0 {
			sum -= right
		} else {
			sum += uniquePathsWithObstaclesStep(obstacleGrid, [2]int{x + 1, y})
		}

	} else {
		flag++
	}

	if y < len(obstacleGrid)-1 {
		down := obstacleGrid[y+1][x]
		if down < 0 {
			sum -= down
		} else {
			sum += uniquePathsWithObstaclesStep(obstacleGrid, [2]int{x, y + 1})
		}
	} else {
		flag++
	}
	if flag == 2 {
		return 1
	}
	if sum == 0 {
		obstacleGrid[y][x] = 1
	} else {
		obstacleGrid[y][x] = -sum
	}

	return sum
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	return uniquePathsWithObstaclesStep(obstacleGrid, [2]int{0, 0})
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1},
	}))
}
