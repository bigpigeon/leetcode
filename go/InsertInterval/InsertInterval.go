/*
57. Insert Interval

Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:
Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].

Example 2:
Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].

This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].
*/
package main

import (
	"fmt"
)

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	m := []Interval{}

	ss, se := 0, len(intervals)
	for ss < se {
		half := (se-ss)/2 + ss
		if intervals[half].Start <= newInterval.Start {
			if intervals[half].End >= newInterval.Start {
				se, ss = half, half
				break
			}
			ss = half + 1
		} else {
			se = half
		}
	}
	fmt.Println(ss, se)
	return m
}

func main() {
	a := []Interval{{1, 3}, {8, 10}, {15, 18}}
	//	sort.Sort(IntervalList(a))
	fmt.Println(insert(a, Interval{6, 14}))

}
