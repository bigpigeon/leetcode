package datatype

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintListNode(node *ListNode) {
	for node != nil {
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func ListToListNode(l []int) *ListNode {
	if len(l) == 0 {
		return nil
	}
	root := &ListNode{l[0], nil}
	curr := root
	for i := 1; i < len(l); i++ {
		curr.Next = &ListNode{l[i], nil}
		curr = curr.Next
	}
	return root
}
