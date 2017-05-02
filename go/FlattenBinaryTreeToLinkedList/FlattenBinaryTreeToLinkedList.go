/*
114. Flatten Binary Tree to Linked List

Given a binary tree, flatten it to a linked list in-place.

For example,
Given

         1
        / \
       2   5
      / \   \
     3   4   6

The flattened tree should look like:


   1
    \
     2
      \
       3
        \
         4
          \
           5
            \
             6

Hints:
If you notice carefully in the flattened tree, each node's right child points to the next node of a pre-order traversal.
*/
package main

import (
	. "github.com/bigpigeon/leetcode/go/datatype"
)

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		//reach to tree tail
		head, tail := root.Left, root.Left
		for tail.Right != nil {
			tail = tail.Right
		}
		root.Left = nil
		root.Right, tail.Right = head, root.Right
	}
}

func main() {
	t1 := ListToTreeNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	flatten(t1)
	PrintTreeNode(t1)
	t2 := ListToTreeNode([]int{1, 2, 5, 3, 4, 6})
	flatten(t2)
	PrintTreeNode(t2)
}
