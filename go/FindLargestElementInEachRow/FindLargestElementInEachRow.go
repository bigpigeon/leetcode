/*
515. Find Largest Element in Each Row

You need to find the largest element in each row of a Binary Tree.

Example:

Input:

          1
         / \
        2   3
       / \   \
      5   3   9

Output: [1, 3, 9]
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

func findValueMostElement(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	largests := []int{root.Val}
	parentCache := []*TreeNode{root}

	for len(parentCache) != 0 {
		nextParentCache := []*TreeNode{}
		var currLargest *TreeNode
		for _, p := range parentCache {
			if p.Left != nil {
				nextParentCache = append(nextParentCache, p.Left)
				// find left most element
				if currLargest == nil || currLargest.Val < p.Left.Val {
					currLargest = p.Left
				}
			}
			if p.Right != nil {
				nextParentCache = append(nextParentCache, p.Right)
				if currLargest == nil || currLargest.Val < p.Right.Val {
					currLargest = p.Right
				}
			}

		}
		if currLargest != nil {
			largests = append(largests, currLargest.Val)
		}
		parentCache = nextParentCache
	}
	return largests
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{5, nil, nil},
			Right: &TreeNode{3, nil, nil},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{9, nil, nil},
		},
	}
	fmt.Println(findValueMostElement(root))
}
