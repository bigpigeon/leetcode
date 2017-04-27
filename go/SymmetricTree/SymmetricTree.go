/*
101. Symmetric Tree

Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3

But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Note:
Bonus points if you could solve it both recursively and iteratively.
*/
package main

import (
	"fmt"

	. "github.com/bigpigeon/leetcode/go/datatype"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	leftStack := []*TreeNode{}
	rightStack := []*TreeNode{}
	stackLen := 0
	for (left != nil || stackLen != 0) && (right != nil || stackLen != 0) {
		if left != nil {
			if right == nil || right.Val != left.Val {
				return false
			}
			leftStack = append(leftStack, left)
			left = left.Left
			rightStack = append(rightStack, right)
			right = right.Right
			stackLen++
		} else {
			if right != nil {
				return false
			}
			stackLen--
			left = leftStack[stackLen].Right
			leftStack = leftStack[:stackLen]
			right = rightStack[stackLen].Left
			rightStack = rightStack[:stackLen]
		}
	}
	return left == right
}

func main() {
	t1 := ListToTreeNode([]int{1, 2, 2, 3, 4, 3, 3})
	fmt.Println(isSymmetric(t1))
}
