/*
124. Binary Tree Maximum Path Sum

Given a binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

For example:
Given the below binary tree,

       1
      / \
     2   3

Return 6.
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func max(v ...int) int {
	m := v[0]
	for i := 1; i < len(v); i++ {
		if m < v[i] {
			m = v[i]
		}
	}
	return m
}

func _maxPathSum(root *TreeNode) (maxShortPath, maxLongPath int) {
	if root.Right == nil && root.Left == nil {
		return root.Val, root.Val
	}
	if root.Right == nil {
		lMaxShortPath, lMaxLongPath := _maxPathSum(root.Left)
		maxShortPath = max(lMaxShortPath, 0) + root.Val
		maxLongPath = max(lMaxLongPath, maxShortPath)
	} else if root.Left == nil {
		rMaxShortPath, rMaxLongPath := _maxPathSum(root.Right)
		maxShortPath = max(rMaxShortPath, 0) + root.Val
		maxLongPath = max(rMaxLongPath, maxShortPath)
	} else {
		lMaxShortPath, lMaxLongPath := _maxPathSum(root.Left)
		rMaxShortPath, rMaxLongPath := _maxPathSum(root.Right)
		maxShortPath = max(lMaxShortPath, rMaxShortPath, 0) + root.Val
		maxLongPath = max(lMaxLongPath, rMaxLongPath, max(lMaxShortPath, 0)+max(rMaxShortPath, 0)+root.Val)
	}

	return
}

func maxPathSum(root *TreeNode) int {
	_, result := _maxPathSum(root)
	return result
}

func main() {
	//	tree := ListToTreeNode([]int{1, 2, 3, 4, 5, 11, 12, 14, 0, 5})

	//	fmt.Println(maxPathSum(tree))
	//	fmt.Println(maxPathSum(&TreeNode{-3, nil, nil}))
	tree2 := ListToTreeNode([]int{2, -1})
	fmt.Println(maxPathSum(tree2))
}
