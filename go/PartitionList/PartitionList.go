/*
86. Partition List

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given 1->4->3->2->5->2 and x = 3,
return 1->2->2->4->3->5.
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

func printNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func ListToNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	pre := head
	for i := 1; i < len(nums); i++ {
		pre.Next = &ListNode{nums[i], nil}
		pre = pre.Next
	}
	return head
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	//rotate head when head value larger than x
	if head.Val >= x {
		curr := head

		for curr.Next != nil && curr.Next.Val >= x {
			curr = curr.Next
		}
		if curr.Next == nil {
			return head
		}
		largerLast := curr
		curr = curr.Next
		for curr.Next != nil && curr.Next.Val < x {
			curr = curr.Next
		}
		smallLast := curr

		smallLast.Next, largerLast.Next, head = head, smallLast.Next, largerLast.Next
		smallLast.Next = partition(smallLast.Next, x)
	} else {
		curr := head
		for curr.Next != nil && curr.Next.Val < x {
			curr = curr.Next
		}
		curr.Next = partition(curr.Next, x)
	}

	return head
}

func main() {
	//	printNode(partition(ListToNode([]int{4, 8, 1, 2, 3, 5, 6}), 4))
	printNode(partition(ListToNode([]int{4, 1, 2, 3, 5, 6, 4, 3, 2, 1}), 4))
	printNode(partition(ListToNode([]int{4, 1, 2, 3}), 4))
	printNode(partition(ListToNode([]int{1, 2, 3, 4, 3, 2, 1}), 4))
}
