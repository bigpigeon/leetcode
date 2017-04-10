/*
98. Validate Binary Search Tree

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:

+-------+
|    2  |
|   / \ |
|  1   3|
+-------+

Binary tree [2,1,3], return true.
Example 2:

+-------+
|    1  |
|   / \ |
|  2   3|
+-------+

Binary tree [1,2,3], return false.
*/
package main

import (
	"fmt"
	"strings"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
			if nums[i] != -1 {
				node.Left = &TreeNode{nums[i], nil, nil}
				currNodes = append(currNodes, node.Left)
			}
			i++
			if i >= len(nums) {
				return root
			}
			if nums[i] != -1 {
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

func printNode(trees ...*TreeNode) {
	for _, tree := range trees {
		printNodeIndent(tree, 0)
	}
}

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	curr := root
	var topLeft, topRight *TreeNode
	for len(stack) != 0 || curr != nil {
		if curr != nil {
			//			fmt.Println(topLeft, topRight, curr.Val)
			if topLeft != nil {
				if curr.Val <= topLeft.Val {
					return false
				}
			}
			if topRight != nil {
				if curr.Val >= topRight.Val {
					return false
				}
			}
			stack = append(stack, curr)
			topRight = curr
			curr = curr.Left
		} else {
			topLeft = stack[len(stack)-1]
			if topLeft == topRight {
				if len(stack) > 1 {
					topRight = stack[len(stack)-2]
				} else {
					topRight = nil
				}
			}
			curr = topLeft.Right
			stack = stack[:len(stack)-1]
		}
	}
	return true
}

func main() {
	//	fmt.Println(isValidBST(listToNode([]int{1, 1})))
	//	fmt.Println(isValidBST(listToNode([]int{1, 0, 2, 3})))
	fmt.Println(isValidBST(listToNode([]int{3, 1, 5, 0, 2, 4, 6, -1, -1, -1, 3})))

}
