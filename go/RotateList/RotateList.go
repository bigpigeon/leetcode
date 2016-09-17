/*
61. Rotate List

Given a list, rotate the list to the right by k places, where k is non-negative.

For example:
Given 1->2->3->4->5->NULL and k = 2,
return 4->5->1->2->3->NULL.
*/
package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func printNode(listNodes ...*ListNode) {
	for _, l := range listNodes {
		for ; l != nil; l = l.Next {
			fmt.Printf("%d->", l.Val)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	r := head
	it := head
	nodeLen := 0
	for it := head; it != nil; it = it.Next {
		nodeLen++
	}
	k = k % nodeLen
	for ; k > 0; k-- {
		if it.Next == nil {
			return head
		}
		it = it.Next
	}
	jt := it
	for ; jt.Next != nil; jt = jt.Next {
		r = r.Next
	}
	jt.Next = head
	head = r.Next
	r.Next = nil
	return head
}

func main() {
	printNode(rotateRight(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	printNode(rotateRight(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 0))
	printNode(rotateRight(&ListNode{1, &ListNode{2, nil}}, 0))
	//	printNode(rotateRight(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 100))
	//	printNode(rotateRight(nil, 4))
}
