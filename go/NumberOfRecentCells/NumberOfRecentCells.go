/*
933. Number of Recent Calls
User Accepted: 1899
User Tried: 2069
Total Accepted: 1934
Total Submissions: 3399
Difficulty: Easy
Write a class RecentCounter to count recent requests.

It has only one method: ping(int t), where t represents some time in milliseconds.

Return the number of pings that have been made from 3000 milliseconds ago until now.

Any ping with time in [t - 3000, t] will count, including the current ping.

It is guaranteed that every call to ping uses a strictly larger value of t than before.



Example 1:

Input: inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]
Output: [null,1,2,3,3]


Note:

Each test case will have at most 10000 calls to ping.
Each test case will call ping with strictly increasing values of t.
Each call to ping will have 1 <= t <= 10^9.
*/

package main

import "fmt"

type TimeCount struct {
	Time, Count int
}

type RecentCounter struct {
	Counter []TimeCount
	Count   int
}

func Constructor() RecentCounter {
	return RecentCounter{Counter: []TimeCount{{0, 0}}, Count: 0}
}

func (this *RecentCounter) Ping(t int) int {
	if t == 0 {
		return this.Count
	}
	last := len(this.Counter) - 1
	this.Count++
	if this.Counter[last].Time == t {
		this.Counter[last].Count++
	} else {
		this.Counter = append(this.Counter, TimeCount{t, 1})
	}
	idx := 0
	for i, c := range this.Counter {
		if c.Time >= t-3000 {
			idx = i
			break
		}
		this.Count -= c.Count
	}
	this.Counter = this.Counter[idx:]
	//fmt.Printf("c %#v\n", this.Counter)
	return this.Count
}

func main() {
	counter := Constructor()
	for _, i := range []int{0, 1, 100, 3001, 3002} {
		fmt.Println(counter.Ping(i))
	}
}
