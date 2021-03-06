/*
82. Remove Duplicates from Sorted List II

Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

For example,
Given 1->2->3->3->4->4->5, return 1->2->5.
Given 1->1->1->2->3, return 2->3.
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

func printNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func deleteDuplicates(head *ListNode) *ListNode {

	//find not duplicate head
	for head != nil && head.Next != nil {
		//this is duplicate element, remove it
		if head.Next.Val == head.Val {
			deplicateVal := head.Val
			for head.Next != nil && head.Next.Val == deplicateVal {
				head = head.Next
			}
			head = head.Next
		} else {
			break
		}
	}
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)

	return head
}

func main() {
	printNode(deleteDuplicates(ListToNode([]int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 8, 8, 9})))
	printNode(deleteDuplicates(ListToNode([]int{})))
	printNode(deleteDuplicates(ListToNode([]int{1})))
}
