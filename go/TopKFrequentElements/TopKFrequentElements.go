/*
347. Top K Frequent Elements

Given a non-empty array of integers, return the k most frequent elements.

For example,
Given [1,1,1,2,2,3] and k = 2, return [1,2].

Note:
1. You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
2. Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/
package main

import (
	"fmt"
	"sort"
)

type FrequentSlice [][2]int

func (p FrequentSlice) Len() int {
	return len(p)
}

func (p FrequentSlice) Less(i, j int) bool {
	return p[i][0] > p[j][0]
}

func (p FrequentSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{} //data struct => map[value]count
	for _, d := range nums {
		m[d]++
	}
	n := make([][2]int, 0, len(m)) // re construct data struct => [[count, value]]
	for k, v := range m {
		n = append(n, [2]int{v, k})
	}
	sort.Sort(FrequentSlice(n))
	frequent := make([]int, 0, k)
	for _, v := range n[:k] {
		frequent = append(frequent, v[1])
	}
	return frequent
}

func main() {
	frequent := topKFrequent([]int{1, 1, 2, 2, 6, 3, 1, 6, 5, 6}, 2)
	fmt.Println(frequent)
}
