/*
24. Swap Nodes in Pairs

Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.
1->2->3
2->3 1->2
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next, next.Next = next.Next, head
	head, next = next, head

	next.Next = swapPairs(next.Next)
	return head
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, nil}}}}
	PrintListNode(swapPairs(l1))
	l2 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, nil}}}}}
	PrintListNode(swapPairs(l2))
	var l3 *ListNode
	PrintListNode(swapPairs(l3))
	l4 := &ListNode{1, nil}
	PrintListNode(swapPairs(l4))
}
