/*
134. Gas Station
Medium

There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique



Example 1:

Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.

Example 2:

Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 3 + 3 = 3
You cannot travel back to station 2, as it requires 4 unit of gas but you only have 3.
Therefore, you can't travel around the circuit once no matter where you start.



Constraints:

    gas.length == n
    cost.length == n
    1 <= n <= 10**4
    0 <= gas[i], cost[i] <= 10**4


*/

package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// resolution:
// create minCost list ,minCost[i]=min(minCost[i-1]+gas[i]-cost[i], gas[i]-cost[i])
// iterator the minCost list ,if have less than zero value, return -1,other way,
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return 0
	}
	c := len(gas) // length of gas or cost
	idx := 0
	idxSum := 0
	gasSum := 0
	for i := 0; i < c; i++ {
		idxSum += gas[i] - cost[i]
		gasSum += gas[i] - cost[i]
		if idxSum < 0 {
			idx = i + 1
			idxSum = 0
		}

	}
	if gasSum < 0 {
		return -1
	}
	return idx
}

func main() {
	idx := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	fmt.Println(idx)
	idx = canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})
	fmt.Println(idx)
	idx = canCompleteCircuit([]int{5, 8, 2, 8}, []int{6, 5, 6, 6})
	fmt.Println(idx)
}
