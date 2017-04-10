/*
513. Find Left Most Element

Given a binary tree, find the left most element in the last row of the tree.

Example 1:

Input:

    2
   / \
  1   3

Output:
1

Example 2:

Input:

        1
       / \
      2   3
     /   / \
    4   5   6
       /
      7

Output:
7
*/
package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeftMostNode(root *TreeNode) int {
	leftMost := root
	parentCache := []*TreeNode{root}

	for len(parentCache) != 0 {
		nextParentCache := []*TreeNode{}
		var currLeftMost *TreeNode

		for _, p := range parentCache {
			if p.Left != nil {
				nextParentCache = append(nextParentCache, p.Left)
				// find left most element
				if currLeftMost == nil {
					currLeftMost = p.Left
				}
			}
			if p.Right != nil {
				nextParentCache = append(nextParentCache, p.Right)
				if currLeftMost == nil {
					currLeftMost = p.Right
				}
			}

		}
		if currLeftMost != nil {
			leftMost = currLeftMost
		}
		parentCache = nextParentCache
	}
	return leftMost.Val
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, nil, nil},
		Right: &TreeNode{3, nil, nil},
	}
	fmt.Println(findLeftMostNode(root))
}
