/*
110. Balanced Binary Tree

Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// if not balanced ,result value was -1
func BalancedAndDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := BalancedAndDepth(root.Left)
	rightDepth := BalancedAndDepth(root.Right)
	//	fmt.Println(leftDepth, rightDepth, root.Left, root.Right)
	if leftDepth == -1 || rightDepth == -1 {
		return -1
	}
	if abs(leftDepth-rightDepth) > 1 {
		return -1
	}
	return max(leftDepth, rightDepth) + 1
}

func isBalanced(root *TreeNode) bool {
	//	fmt.Println(BalancedAndDepth(root))
	return BalancedAndDepth(root) != -1
}

func main() {
	t1 := ListToTreeNode([]int{1, 2, 3, 4, 5, 6, 8, 9, 10})
	fmt.Println(isBalanced(t1))
	t2 := ListToTreeNode([]int{1, 2, 3, 4, 5, 6, 8, 9, 0, 10})
	fmt.Println(isBalanced(t2))
	t3 := ListToTreeNode([]int{1, 2, 3, 4, 5, 6, 8, 9, 0, 10, 0, 2, 0, 3, 2, 4})
	fmt.Println(isBalanced(t3))
	t4 := ListToTreeNode([]int{1, 0, 2, 0, 3})
	fmt.Println(isBalanced(t4))
	t5 := ListToTreeNode([]int{3, 1, 0, 0, 2})
	fmt.Println(isBalanced(t5))
}
