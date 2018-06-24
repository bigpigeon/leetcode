/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := 0
	profitCache := make([]int, len(prices))
	minCache := make([]int, len(prices))
	for i, min := 1, 0; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		} else {
			profitCache[i] = prices[i] - prices[min]
			if profitCache[i] > profitCache[max] {
				max = i
			}
		}
		minCache[i] = min
	}
	// find left
	sec := 0
	for i := 1; i < minCache[max]; i++ {
		if profitCache[i] > sec {
			sec = profitCache[i]
		}
	}
	// find right
	secMin := max + 1
	for i := secMin + 1; i < len(prices); i++ {
		if prices[i] < prices[secMin] {
			secMin = i
		} else if prices[i]-prices[secMin] > sec {
			sec = prices[i] - prices[secMin]
		}
	}
	// find middle
	secMax := minCache[max] + 1
	for i := secMax + 1; i < max; i++ {
		if prices[i] > prices[secMax] {
			secMax = i
		} else if prices[secMax]-prices[i] > sec {
			sec = prices[secMax] - prices[i]
		}
	}

	return prices[max] - prices[minCache[max]] + sec
}

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
	fmt.Println(maxProfit([]int{6, 1, 3, 2, 4, 7}))
	fmt.Println(maxProfit([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}))
}
