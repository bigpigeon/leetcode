/*
92. Reverse Linked List II

Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,

return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.

step
1->2->3->4->5->NULL

1->2<-3 4->5->NULL

1->2<-3<-4 5->NULL
  |      |
  |______|

1->4->3->2 5->NULL
1->4->3->2->5->NULL
*/
package main

import (
	"fmt"
)

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateNode(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	return &ListNode{list[0], CreateNode(list[1:])}

}

func PrintNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var preMNode, mNode *ListNode
	if m == 1 {
		preMNode = nil
		mNode = head
	} else {
		preMNode = head
		for i := 2; i < m; i++ {
			preMNode = preMNode.Next
		}
		mNode = preMNode.Next
	}
	nNode := mNode
	suffNNode := mNode.Next
	/*
		e.g 1->2->3->4->5->NULL
		    1->2<-3  4->5->NULL
			1->2<-3<-4  5->NULL
	*/
	for i := m + 1; i <= n; i++ {
		next := suffNNode.Next
		suffNNode.Next = nNode
		nNode = suffNNode
		suffNNode = next
	}
	if preMNode != nil {
		preMNode.Next = nNode
	} else {
		head = nNode
	}
	mNode.Next = suffNNode
	return head
}

func main() {
	PrintNode(reverseBetween(CreateNode([]int{1, 2, 3, 4, 5}), 2, 4))
	PrintNode(reverseBetween(CreateNode([]int{1, 2, 3, 4, 5}), 2, 5))
	PrintNode(reverseBetween(CreateNode([]int{1, 2, 3, 4, 5}), 1, 3))
	PrintNode(reverseBetween(CreateNode([]int{1, 2, 3, 4, 5}), 1, 5))
}
