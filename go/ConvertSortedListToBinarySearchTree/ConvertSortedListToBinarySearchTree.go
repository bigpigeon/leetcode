/*
109. Convert Sorted List to Binary Search Tree

Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
*/
package main

import (
	. "github.com/bigpigeon/leetcode/go/datatype"
)

func toBST(head, tail *ListNode) *TreeNode {
	fast := head
	slow := head
	if fast == tail {
		return nil
	}
	for fast.Next != tail && fast.Next.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return &TreeNode{slow.Val, toBST(head, slow), toBST(slow.Next, tail)}
}

func sortedListToBST(head *ListNode) *TreeNode {
	return toBST(head, nil)
}

func main() {
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	t1 := sortedListToBST(ListToListNode(a1))
	PrintTreeNode(t1)
	a2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t2 := sortedListToBST(ListToListNode(a2))
	PrintTreeNode(t2)
}
