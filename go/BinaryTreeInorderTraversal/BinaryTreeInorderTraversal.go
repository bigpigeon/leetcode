/*
94. Binary Tree Inorder Traversal

Given a binary tree, return the inorder traversal of its nodes' values.

For example:
Given binary tree [1,null,2,3],

+-------+
|  1    |
|   \   |
|    2  |
|   /   |
|  3    |
+-------+

return [1,3,2].

Note: Recursive solution is trivial, could you do it iteratively?
*/

package main

import (
	"fmt"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printNodeIndent(root *TreeNode, indent int) {
	if root != nil {
		if indent == 0 {
			fmt.Println(root.Val)
		} else {
			fmt.Println(strings.Repeat(" ", indent-2)+"|-", root.Val)
		}
		printNodeIndent(root.Left, indent+2)
		printNodeIndent(root.Right, indent+2)
	}
}

func printNode(root *TreeNode) {
	printNodeIndent(root, 0)
}

func listToNode(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{nums[0], nil, nil}
	preNodes := []*TreeNode{root}
	currNodes := []*TreeNode{}
	i := 1
	for i < len(nums) {
		for _, node := range preNodes {
			if nums[i] != 0 {
				node.Left = &TreeNode{nums[i], nil, nil}
				currNodes = append(currNodes, node.Left)
			}
			i++
			if i >= len(nums) {
				return root
			}
			if nums[i] != 0 {
				node.Right = &TreeNode{nums[i], nil, nil}
				currNodes = append(currNodes, node.Right)
			}
			i++
			if i >= len(nums) {
				return root
			}
		}
		preNodes, currNodes = currNodes, []*TreeNode{}
	}
	return root
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{}
	for curr := root; curr != nil || len(stack) != 0; {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		result = append(result, curr.Val)
		curr = curr.Right
	}
	return result
}

func main() {
	fmt.Println(inorderTraversal(listToNode([]int{1, 0, 2, 3, 5, 6, 7})))
	//	fmt.Println(inorderTraversal(listToNode([]int{1, 3, 5, 0, 7, 9, 10, 11, 12, 13, 14, 15, 16})))
}
