/*
99. Recover Binary Search Tree

Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

Note:
A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?


solution: Morris Traversal
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func recoverTree(root *TreeNode) {
	curr := root
	var prev *TreeNode
	var firstElement, secondElement *TreeNode
	var parentElement *TreeNode
	for curr != nil {
		if curr.Left == nil {
			fmt.Println(curr.Val)
			if parentElement != nil && parentElement.Val > curr.Val {
				if firstElement == nil {
					firstElement = parentElement
				}
				secondElement = curr
			}
			parentElement = curr

			curr = curr.Right
		} else {
			prev = curr.Left
			for prev.Right != nil && prev.Right != curr {
				prev = prev.Right
			}

			if prev.Right == nil {
				prev.Right = curr
				curr = curr.Left
			} else {
				prev.Right = nil
				fmt.Println(curr.Val)
				if parentElement != nil && parentElement.Val > curr.Val {
					if firstElement == nil {
						firstElement = parentElement
					}
					secondElement = curr
				}
				parentElement = curr

				curr = curr.Right
			}
		}
	}
	if firstElement != nil && secondElement != nil {
		firstElement.Val, secondElement.Val = secondElement.Val, firstElement.Val
	}
}

func main() {
	//	{
	//		node := ListToTreeNode([]int{1, 2})
	//		recoverTree(node)
	//		PrintTreeNode(node)
	//	}
	//	{
	//		node := ListToTreeNode([]int{2, 1, 3})
	//		recoverTree(node)
	//		PrintTreeNode(node)
	//	}
	{
		node := ListToTreeNode([]int{1, 2, 3})
		recoverTree(node)
		PrintTreeNode(node)
	}
	{
		node := ListToTreeNode([]int{3, 0, 2, 0, 1})
		recoverTree(node)
		PrintTreeNode(node)
	}
}
