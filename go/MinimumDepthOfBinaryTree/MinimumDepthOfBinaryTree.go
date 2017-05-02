/*
111. Minimum Depth of Binary Tree

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	//	findleft := false
	layer := []*TreeNode{root}
	for len(layer) != 0 {
		depth++
		newlayer := []*TreeNode{}
		for _, l := range layer {
			if l.Left == nil && l.Right == nil {
				return depth
			}
			if l.Left != nil {
				newlayer = append(newlayer, l.Left)
			}
			if l.Right != nil {
				newlayer = append(newlayer, l.Right)
			}
		}
		layer = newlayer
	}
	return depth
}

func main() {
	const null = 0
	t1 := ListToTreeNode([]int{1, 2})
	fmt.Println(minDepth(t1))
	t2 := ListToTreeNode([]int{1, 2, 3, 1, null, 4, null, 4})
	fmt.Println(minDepth(t2))
	t3 := ListToTreeNode([]int{1, null, 2})
	fmt.Println(minDepth(t3))
}
