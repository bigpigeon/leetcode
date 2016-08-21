/*
19. Remove Nth Node From End of List

Given a linked list, remove the nth node from the end of list and return its head.

For example,
+-----------------------------------------------------------------------------------+
|   Given linked list: 1->2->3->4->5, and n = 2.                                    |
|                                                                                   |
|   After removing the second node from the end, the linked list becomes 1->2->3->5.|
+-----------------------------------------------------------------------------------+

Note:
Given n will always be valid.
Try to do this in one pass.
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

func printNode(n *ListNode) {
	for n != nil {
		fmt.Printf("%d->", n.Val)
		n = n.Next
	}
	fmt.Println()
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nListNode := []*ListNode{head}
	node := head.Next
	for node != nil {
		nListNode = append(nListNode, node)
		node = node.Next
	}
	delIndex := len(nListNode) - n
	if delIndex < 0 || delIndex >= len(nListNode) {
		return head
	}

	if delIndex == 0 { // remove index at the head
		head = head.Next
	} else {
		nListNode[delIndex-1].Next = nListNode[delIndex].Next
	}
	return head
}

func main() {
	n1 := &ListNode{1, nil}
	printNode(removeNthFromEnd(n1, 1))
	for i := 0; i <= 6; i++ {
		fmt.Println("remove index", i)
		n1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		printNode(removeNthFromEnd(n1, i))
	}

}
