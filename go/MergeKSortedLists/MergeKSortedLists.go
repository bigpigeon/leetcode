/*
23. Merge k Sorted Lists
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

best solution:
merge as a pair witch length is closely

a example ListNode array:
+------------------------------------------------------------------------------------+
| [ l1(len:3), l2(len:3), l3(len:5), l4(len:7), l5(len:7), l6(len:10), l7(len:15) ]  |
|                                                                                    |
| l1 # l2 => #l3 => #l6 =#                                                           |
|                        #=> #l7                                                     |
|             l4 # l5 ===#                                                           |
+------------------------------------------------------------------------------------+

now solution:
merge as a pair witch position is closely
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

func PrintListNode(l *ListNode) {
	for l != nil {
		fmt.Printf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	l := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			l.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		l = l.Next
	}
	for l1 != nil {
		l.Next = &ListNode{l1.Val, nil}
		l1 = l1.Next
		l = l.Next
	}
	for l2 != nil {
		l.Next = &ListNode{l2.Val, nil}
		l2 = l2.Next
		l = l.Next
	}
	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mergeList := []*ListNode{}

	for i := 1; i < len(lists); i += 2 {
		mergeList = append(mergeList, mergeTwoLists(lists[i-1], lists[i]))
	}
	merge := mergeKLists(mergeList)
	// merge the lastest ListNode when the lists lenght is odd number
	if len(lists)%2 == 1 {
		merge = mergeTwoLists(merge, lists[len(lists)-1])
	}
	return merge
}

func main() {
	l1 := &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{9, nil}}}}
	l2 := &ListNode{2, &ListNode{4, &ListNode{6, &ListNode{8, nil}}}}
	l3 := &ListNode{7, &ListNode{11, nil}}
	PrintListNode(mergeKLists([]*ListNode{l1, l2, l3}))
	PrintListNode(mergeKLists([]*ListNode{}))
	PrintListNode(mergeKLists([]*ListNode{l1}))
}
