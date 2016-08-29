/*
21. Merge Two Sorted Lists
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ml := new(ListNode)
	curr := ml

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		curr = curr.Next
	}
	for l1 != nil {
		curr.Next = &ListNode{l1.Val, nil}
		l1 = l1.Next
		curr = curr.Next
	}
	for l2 != nil {
		curr.Next = &ListNode{l2.Val, nil}
		l2 = l2.Next
		curr = curr.Next
	}
	return ml.Next
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{7, &ListNode{9, nil}}}}}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, &ListNode{10, nil}}}}}
	PrintListNode(mergeTwoLists(l1, l2))
	PrintListNode(mergeTwoLists(nil, nil))
}
