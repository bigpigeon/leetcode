/*
100. Same Tree

Given two binary trees, write a function to check if they are equal or not.

Two binary trees are considered equal if they are structurally identical and the nodes have the same value.
*/
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pStack := []*TreeNode{}
	qStack := []*TreeNode{}
	for p != nil || len(pStack) != 0 {
		if p != nil {
			if q != nil && q.Val == p.Val {
				qStack = append(qStack, q)
				q = q.Left
			} else {
				return false
			}
			pStack = append(pStack, p)
			p = p.Left
		} else {
			if q != nil {
				return false
			}
			p = pStack[len(pStack)-1].Right
			pStack = pStack[:len(pStack)-1]
			q = qStack[len(qStack)-1].Right
			qStack = qStack[:len(qStack)-1]
		}
	}
	return len(qStack) == 0 && q == nil
}

func main() {
	fmt.Println(isSameTree((*TreeNode)(nil), &TreeNode{2, nil, nil}))
}
