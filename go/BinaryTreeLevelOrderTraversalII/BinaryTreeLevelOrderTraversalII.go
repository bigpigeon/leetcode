/*
107. Binary Tree Level Order Traversal II

Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

return its bottom-up level order traversal as:

+---------+
|[        |
|  [15,7],|
|  [9,20],|
|  [3]    |
|]        |
+---------+
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func CombineListTail(l1, l2 [][]int) [][]int {
	var long, short [][]int
	reverse := false
	if len(l1) > len(l2) {
		long, short = l1, l2
	} else {
		long, short = l2, l1
		reverse = true
	}
	delta := len(long) - len(short)
	newList := [][]int{}
	for i := 0; i < delta; i++ {
		newList = append(newList, append([]int{}, long[i]...))
	}
	for j := 0; j < len(short); j++ {
		if reverse {
			newList = append(newList, append([]int{}, short[j]...))
			newList[j+delta] = append(newList[j+delta], long[j+delta]...)
		} else {

			newList = append(newList, append([]int{}, long[j+delta]...))
			newList[j+delta] = append(newList[j+delta], short[j]...)
		}
	}
	return newList
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	l1 := levelOrderBottom(root.Left)
	l2 := levelOrderBottom(root.Right)
	newl := CombineListTail(l1, l2)
	newl = append(newl, []int{root.Val})
	return newl
}

func main() {
	l1 := [][]int{
		[]int{1, 2},
		[]int{3, 4},
		[]int{4, 2},
		[]int{9, 10},
	}
	l2 := [][]int{
		[]int{11, 12},
		[]int{13, 14},
	}
	fmt.Println(CombineListTail(l1, l2))
	fmt.Println(CombineListTail(l2, l1))

	n1 := ListToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(levelOrderBottom(n1))
}
