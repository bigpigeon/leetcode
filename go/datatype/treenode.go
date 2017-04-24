package datatype

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTreeNodeIndent(root *TreeNode, indent int) {
	if root != nil {
		if indent == 0 {
			fmt.Println(root.Val)
		} else {
			fmt.Println(strings.Repeat(" ", indent-2)+"|-", root.Val)
		}
		if root.Left == nil {
			fmt.Println(strings.Repeat(" ", indent) + "|-")
		} else {
			printTreeNodeIndent(root.Left, indent+2)
		}

		printTreeNodeIndent(root.Right, indent+2)
	}
}

func PrintTreeNode(root *TreeNode) {
	printTreeNodeIndent(root, 0)
}

func ListToTreeNode(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{nums[0], nil, nil}
	preNodes := []*TreeNode{root}
	currNodes := []*TreeNode{}
	i := 1
	for i < len(nums) {
		for _, node := range preNodes {
			if nums[i] != 0 {
				node.Left = &TreeNode{nums[i], nil, nil}
				currNodes = append(currNodes, node.Left)
			}
			i++
			if i >= len(nums) {
				return root
			}
			if nums[i] != 0 {
				node.Right = &TreeNode{nums[i], nil, nil}
				currNodes = append(currNodes, node.Right)
			}
			i++
			if i >= len(nums) {
				return root
			}
		}
		preNodes, currNodes = currNodes, []*TreeNode{}
	}
	return root
}
