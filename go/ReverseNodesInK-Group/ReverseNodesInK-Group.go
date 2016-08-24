/*
25. Reverse Nodes in k-Group

Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

You may not alter the values in the nodes, only nodes itself may be changed.

Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
*/
package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}
	node := head
	nodes := make([]*ListNode, 0, k)

	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		nodes = append(nodes, node)
		node = node.Next
	}
	head.Next = reverseKGroup(nodes[k-1].Next, k)
	for i := k - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}

	return nodes[k-1]
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, nil}}}}
	PrintListNode(reverseKGroup(l1, 2))
	l2 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, nil}}}}}
	PrintListNode(reverseKGroup(l2, 3))
	var l3 *ListNode
	PrintListNode(reverseKGroup(l3, 2))
	l4 := &ListNode{1, nil}
	PrintListNode(reverseKGroup(l4, 2))
	l5 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, nil}}}}
	PrintListNode(reverseKGroup(l5, 1))
}
