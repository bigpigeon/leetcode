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
