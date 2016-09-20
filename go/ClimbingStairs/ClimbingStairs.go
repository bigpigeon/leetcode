/*
70. Climbing Stairs

You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
solution:
  Fibonacci
*/
package main

import (
	"fmt"
)

func climbStairs(n int) int {
	cache := []int{0, 1}
	for i := 2; i <= n+1; i++ {
		cache = append(cache, cache[i-1]+cache[i-2])
	}
	return cache[n+1]
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
}
