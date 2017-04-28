/*
105. Construct Binary Tree from Preorder and Inorder Traversal

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

solution :
the preorder list first element must be root

the inorder list first element must be the leftmost element

the following struct was a  preorder list as three depth full tree

[root, root.Left, root.Left.Left, root.Left.Right, root.Right, root.Right.Left, root.Right.Right]

the following struct was a inorder list as three depth full tree

[root.Left.Left, root.Left, root.Left.Right, root, root.Right.Left, root.Right, root.Right.Right]

the easiest solution:
1. find the root value postion with inorder
2. split all left values and right values(preorder split [1:pos+1], [pos+1:], inorder split [:pos], [pos+1:])
3. recursion
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

// return val postion, if not return -1
func findOrderVal(order []int, val int) int {
	mid := len(order) - 1
	for i := mid - 1; i >= 0; i-- {
		if order[i] == val {
			return i
		}
	}
	for i := mid; i < len(order); i++ {
		if order[i] == val {
			return i
		}
	}
	return -1
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// find the root at inorder position
	rootVal := preorder[0]
	root := &TreeNode{rootVal, nil, nil}
	var leftInorder, rightInorder []int
	var leftPreorder, rightPreorder []int

	if i := findOrderVal(inorder, rootVal); i != -1 {
		leftPreorder, rightPreorder = preorder[1:i+1], preorder[i+1:]
		leftInorder, rightInorder = inorder[:i], inorder[i+1:]
	}
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

func main() {
	n1 := ListToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	PrintTreeNode(n1)
	n1Preorder := Preorder(n1)
	n1Inorder := Inorder(n1)
	fmt.Println(n1Preorder)
	fmt.Println(n1Inorder)
	PrintTreeNode(buildTree(n1Preorder, n1Inorder))
}
