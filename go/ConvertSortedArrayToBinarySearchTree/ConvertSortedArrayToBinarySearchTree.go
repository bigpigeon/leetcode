/*
108. Convert Sorted Array to Binary Search Tree

Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
*/
package main

import (
	. "github.com/bigpigeon/leetcode/go/datatype"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	haft := len(nums) / 2
	root := &TreeNode{nums[haft], nil, nil}
	root.Left = sortedArrayToBST(nums[:haft])
	root.Right = sortedArrayToBST(nums[haft+1:])
	return root
}

func main() {
	a1 := []int{1, 3, 5, 7, 9, 11}
	PrintTreeNode(sortedArrayToBST(a1))
	a2 := []int{1, 2, 3, 4, 5, 6}
	PrintTreeNode(sortedArrayToBST(a2))
}
