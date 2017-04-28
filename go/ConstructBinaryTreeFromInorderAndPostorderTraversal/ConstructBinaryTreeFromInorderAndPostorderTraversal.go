/*
106. Construct Binary Tree from Inorder and Postorder Traversal

Given inorder and postorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

solution:
similar preorder and inorder traversal problem
just root position was at the last element in postorder

the following struct was a inorder list as three depth full tree

[root.Left.Left, root.Left, root.Left.Right, root, root.Right.Left, root.Right, root.Right.Right]

the following struct was a postorder list at three depth full tree

[root.Left.Left, root.Left.Right, root.Left, root.Right.Left, root.Right.Right, root.Right, root]
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	// find the root at inorder position
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{rootVal, nil, nil}
	var leftInorder, rightInorder []int
	var leftPostorder, rightPostorder []int

	if i := findOrderVal(inorder, rootVal); i != -1 {
		leftInorder, rightInorder = inorder[:i], inorder[i+1:]
		leftPostorder, rightPostorder = postorder[0:i], postorder[i:len(postorder)-1]
	}
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return root
}

func main() {
	n1 := ListToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	PrintTreeNode(n1)
	n1Inorder := Inorder(n1)
	n1Postorder := Postorder(n1)
	fmt.Println(n1Inorder)
	fmt.Println(n1Postorder)
	PrintTreeNode(buildTree(n1Inorder, n1Postorder))
}
