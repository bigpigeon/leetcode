/*
56. Merge Intervals

Given a collection of intervals, merge all overlapping intervals.

For example,
Given [1,3],[2,6],[8,10],[15,18],
return [1,6],[8,10],[15,18].
*/
package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for an interval.
 */
type Interval struct {
	Start int
	End   int
}

type IntervalList []Interval

func (l IntervalList) Len() int {
	return len(l)
}

func (l IntervalList) Less(i, j int) bool {
	return l[i].Start < l[j].Start
}

func (l IntervalList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Sort(IntervalList(intervals))
	m := []Interval{}
	last := intervals[0]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if curr.Start <= last.End {
			last.End = max(curr.End, last.End)
		} else {
			m = append(m, last)
			last = curr
		}
	}
	m = append(m, last)
	return m
}

func main() {
	a := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {9, 16}}
	//	sort.Sort(IntervalList(a))
	fmt.Println(merge(a))
	fmt.Println(merge([]Interval{}))
}
