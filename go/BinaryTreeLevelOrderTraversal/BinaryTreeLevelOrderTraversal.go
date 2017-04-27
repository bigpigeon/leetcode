/*
102. Binary Tree Level Order Traversal

Total Accepted: 163948
Total Submissions: 426709
Difficulty: Medium
Contributor: LeetCode
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its level order traversal as:

+---------+
|[        |
|  [3],   |
|  [9,20],|
|  [15,7] |
|]        |
+---------+

*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func levelOrder(root *TreeNode) [][]int {
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
			order[currLevel] = append(order[currLevel], root.Val)
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
	fmt.Println(levelOrder(t1))
}
