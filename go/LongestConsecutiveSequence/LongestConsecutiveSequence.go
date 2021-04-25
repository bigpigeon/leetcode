/*
128. Longest Consecutive Sequence
Hard

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9



Constraints:

    0 <= nums.length <= 104
    -109 <= nums[i] <= 109


Follow up: Could you implement the O(n) solution?
*/
package main

import "fmt"

type Node struct {
	Start, End int

	Pre, Back *Node
}

func (n *Node) Insert(v int) *Node {
	if v == n.Start-1 {
		n.Start = n.Start - 1
		n.TryMergePre()
	}
	if v == n.End+1 {
		n.End = n.End + 1
		n.TryMergeBack()
	}
	if v < n.Start-1 {
		if n.Pre == nil || v > n.Pre.End+1 {
			pre := &Node{Start: v, End: v}
			n.InsertPreNode(pre)
		} else {
			return n.Pre.Insert(v)
		}
	}
	if v > n.End+1 {
		if n.Back == nil || v < n.Back.Start-1 {
			back := &Node{Start: v, End: v}
			n.InsertBackNode(back)
		} else {
			return n.Back.Insert(v)
		}
	}
	return n
}

func (n *Node) InsertBackNode(back *Node) {
	backBack := n.Back
	n.Back = back
	back.Pre = n
	back.Back = backBack
	if backBack != nil {
		backBack.Pre = back
	}
}

func (n *Node) InsertPreNode(pre *Node) {
	prePre := n.Pre
	n.Pre = pre
	pre.Back = n
	pre.Pre = prePre
	if prePre != nil {
		prePre.Back = pre
	}
}

func (n *Node) TryMergePre() {
	if n.Pre != nil && (n.Pre.End == n.Start || n.Pre.End == n.Start-1) { // pre node less one or equal to n.Start
		n.Start = n.Pre.Start
		// remove pre node
		if n.Pre.Pre != nil {
			n.Pre.Pre.Back = n
		}
		n.Pre = n.Pre.Pre
	}
}

func (n *Node) TryMergeBack() {
	if n.Back != nil && (n.Back.Start == n.End || n.Back.Start == n.End+1) { // back node start equal or more one than n.End merge it
		n.End = n.Back.End
		// remove pre node
		if n.Back.Back != nil {
			n.Back.Back.Pre = n
		}
		n.Back = n.Back.Back
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	consecutiveList := &Node{Start: nums[0], End: nums[0]} // represent start1, end1, start2, end2

	for _, v := range nums[1:] {
		consecutiveList = consecutiveList.Insert(v)
	}
	maxLen := consecutiveList.End - consecutiveList.Start + 1
	for node := consecutiveList.Pre; node != nil; node = node.Pre {
		l := node.End - node.Start + 1
		if l > maxLen {
			maxLen = l
		}
	}
	for node := consecutiveList.Back; node != nil; node = node.Back {
		l := node.End - node.Start + 1
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

func main() {
	//fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
