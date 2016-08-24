/*
2. Add Two Numbers

You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lsum := new(ListNode)
	var plsum_front *ListNode
	plsum := lsum
	for ; l1 != nil && l2 != nil; l1, l2, plsum_front, plsum = l1.Next, l2.Next, plsum, plsum.Next {
		sum := l1.Val + l2.Val + plsum.Val
		plsum.Val = sum % 10
		plsum.Next = &ListNode{Val: sum / 10}
	}
	for ; l1 != nil; l1, plsum_front, plsum = l1.Next, plsum, plsum.Next {
		sum := l1.Val + plsum.Val
		plsum.Val = sum % 10
		plsum.Next = &ListNode{Val: sum / 10}
	}
	for ; l2 != nil; l2, plsum_front, plsum = l2.Next, plsum, plsum.Next {
		sum := l2.Val + plsum.Val
		plsum.Val = sum % 10
		plsum.Next = &ListNode{Val: sum / 10}
	}
	//delete zero number with top node
	if plsum_front != nil && plsum.Val == 0 {
		plsum_front.Next = nil
	}
	return lsum

}

func (n *ListNode) Print() {
	for n != nil {
		fmt.Printf("%d->", n.Val)
		n = n.Next
	}
	fmt.Printf("\n")
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(l1, l2).Print()
	l1 = &ListNode{2, &ListNode{9, &ListNode{9, nil}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{9, &ListNode{9, nil}}}}
	addTwoNumbers(l1, l2).Print()
	l1 = &ListNode{0, nil}
	l2 = &ListNode{0, nil}
	addTwoNumbers(l1, l2).Print()
}
