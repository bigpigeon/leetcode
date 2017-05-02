/*
113. Path Sum II

Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
Given the below binary tree and sum = 22,

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1

return
+--------------+
|[             |
|   [5,4,11,2],|
|   [5,8,4,5]  |
|]             |
+--------------+
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return [][]int{[]int{root.Val}}
		}
	}
	fmt.Println(root.Val, sum)
	solution := [][]int{}
	if root.Left != nil {
		ret := pathSum(root.Left, sum-root.Val)
		for _, e := range ret {
			v := append([]int{root.Val}, e...)
			solution = append(solution, v)
		}
	}
	if root.Right != nil {
		ret := pathSum(root.Right, sum-root.Val)
		for _, e := range ret {
			v := append([]int{root.Val}, e...)
			solution = append(solution, v)
		}
	}
	return solution
}

func main() {
	const null = 0
	t1 := ListToTreeNode([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1})
	fmt.Println(pathSum(t1, 22))
}
