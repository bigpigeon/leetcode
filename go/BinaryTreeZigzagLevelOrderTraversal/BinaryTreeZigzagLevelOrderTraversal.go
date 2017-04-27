/*
103. Binary Tree Zigzag Level Order Traversal

Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its zigzag level order traversal as:

[
  [3],
  [20,9],
  [15,7]
]
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	stack := []struct {
		Node  *TreeNode
		Level int
	}{}
	currLevel := 0
	order := [][]int{}

	for root != nil || len(stack) != 0 {
		if root != nil {

			if len(order) <= currLevel {
				order = append(order, []int{})
			}
			if currLevel%2 == 0 {
				order[currLevel] = append(order[currLevel], root.Val)
			} else {
				order[currLevel] = append([]int{root.Val}, order[currLevel]...)
			}
			stack = append(stack, struct {
				Node  *TreeNode
				Level int
			}{root, currLevel})
			root = root.Left
			currLevel++
		} else {
			root = stack[len(stack)-1].Node.Right
			currLevel = stack[len(stack)-1].Level + 1
			stack = stack[:len(stack)-1]
		}
	}
	return order
}

func main() {
	t1 := ListToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(zigzagLevelOrder(t1))
}
