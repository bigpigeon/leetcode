/*
 * Copyright 2018. bigpigeon. All rights reserved.
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 */

package main

import (
	"fmt"
	. "github.com/bigpigeon/leetcode/go/datatype"
	"runtime"
)

func PaveFn(node *TreeNode, level int, fn func(*TreeNode) bool) bool {
	if level != 1 {
		if PaveFn(node.Right, level-1, fn) == false {
			return PaveFn(node.Left, level-1, fn)
		}
		return true
	} else {
		return fn(node)
	}

}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rightmost := root
	var stack []*TreeNode
	for rightmost.Right != nil {
		stack = append(stack, rightmost.Left)
		rightmost = rightmost.Right

	}
	level := len(stack) + 1 // last level
	stack = append(stack, rightmost)
	full := 1 << uint(level)
	var findEnd bool
	PaveFn(root, level, func(node *TreeNode) bool {
		if node.Right != nil {
			findEnd = true
		} else if node.Left != nil {
			findEnd = true
			full--
		} else {
			full -= 2
		}
		return findEnd
	})
	return full + (1<<uint(level) - 1)
}

func main() {
	runtime.GOMAXPROCS(1)
	{
		count := countNodes(ListToTreeNode(MakeRange(1, 1)))
		fmt.Println(count)
	}
	{
		count := countNodes(ListToTreeNode(MakeRange(1, 2)))
		fmt.Println(count)
	}
	{
		count := countNodes(ListToTreeNode(MakeRange(1, 7)))
		fmt.Println(count)
	}
	{
		count := countNodes(ListToTreeNode(MakeRange(1, 8)))
		fmt.Println(count)
	}
	{

		count := countNodes(ListToTreeNode(MakeRange(1, 9)))
		fmt.Println(count)
	}
	{

		count := countNodes(ListToTreeNode(MakeRange(1, 10)))
		fmt.Println(count)
	}
	{

		count := countNodes(ListToTreeNode(MakeRange(1, 1000)))
		fmt.Println(count)
	}
}
