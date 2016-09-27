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
	startVal := newInterval.Start
	endVal := newInterval.End
	//find the newInterval.Start position
	ss, se := 0, len(intervals)

	for ss < se {
		half := (se-ss)/2 + ss
		start, end := intervals[half].Start, intervals[half].End
		if newInterval.Start >= start && newInterval.Start <= end {
			ss, se = half, half
			startVal = start
			break
		} else if newInterval.Start < start {
			se = half
		} else {
			ss = half + 1
		}
	}
	fmt.Println(ss, se)

	if ss < len(intervals) {
		//find the newInterval.End position
		es, ee := ss, len(intervals)
		for es < ee {
			half := (ee-es)/2 + es
			start, end := intervals[half].Start, intervals[half].End
			if newInterval.End >= start && newInterval.End <= end {
				es, ee = half+1, half+1
				endVal = end
				break
			} else if newInterval.End < start {
				ee = half
			} else {
				es = half + 1
			}
		}
		solution := make([]Interval, len(intervals[:ss]))
		copy(solution, intervals[:ss])
		solution = append(solution, Interval{startVal, endVal})
		if es < len(intervals) {
			solution = append(solution, intervals[es:]...)
		}
		return solution
	}
	return append(intervals, Interval{startVal, endVal})

}

func main() {
	a := []Interval{{1, 3}, {8, 10}, {15, 18}}

	//	sort.Sort(IntervalList(a))
	fmt.Println(insert(a, Interval{6, 14}))
	fmt.Println(insert([]Interval{{1, 3}, {4, 5}, {7, 10}, {15, 18}}, Interval{6, 15}))
	fmt.Println(insert([]Interval{{1, 3}}, Interval{6, 14}))

}
