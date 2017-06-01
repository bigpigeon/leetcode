/*
122. Best Time to Buy and Sell Stock II

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times). However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
*/
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	currProfit := 0
	sumProfit := 0
	for i := 1; i < len(prices); i++ {
		currProfit = max(prices[i]-min, currProfit)
		if prices[i] < prices[i-1] {
			sumProfit += currProfit
			min = prices[i]
			currProfit = 0
		}
	}
	sumProfit += currProfit
	return sumProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 4, 2, 5, 1, 3}))
}
