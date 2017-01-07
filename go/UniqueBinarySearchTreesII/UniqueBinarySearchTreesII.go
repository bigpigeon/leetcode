/*
95. Unique Binary Search Trees II

Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1...n.

For example,
Given n = 3, your program should return all 5 unique BST's shown below.
+-------------------------------------------+
|   1         3     3      2      1         |
|    \       /     /      / \      \        |
|     3     2     1      1   3      2       |
|    /     /       \                 \      |
|   2     1         2                 3     |
+-------------------------------------------+
*/
package main

import (
	"fmt"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Clone(n *TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	return &TreeNode{n.Val, Clone(n.Left), Clone(n.Right)}
}

func printNodeIndent(root *TreeNode, indent int) {
	if root != nil {
		if indent == 0 {
			fmt.Println(root.Val)
		} else {
			fmt.Println(strings.Repeat(" ", indent-2)+"|-", root.Val)
		}
		printNodeIndent(root.Left, indent+2)
		printNodeIndent(root.Right, indent+2)
	}
}

func printNode(trees ...*TreeNode) {
	for _, tree := range trees {
		printNodeIndent(tree, 0)
	}
}

func generateTrees(n int) []*TreeNode {
	// hash[start][length] start mean start num, length means end -start
	hash := map[int]map[int][]*TreeNode{}
	for i := 1; i <= n; i++ {
		hash[i] = map[int][]*TreeNode{}
		hash[i][1] = []*TreeNode{{i, nil, nil}}
	}
	// full
	hash[n+1] = map[int][]*TreeNode{}

	for length := 2; length <= n; length++ {
		for i := 1; i <= n-length+1; i++ {

			result := []*TreeNode{}
			for j := i; j < i+length; j++ {
				fmt.Println(i, j-1, ":", j, ":", j+1, i+length-j-1)
				small := hash[i][j-i]
				big := hash[j+1][i+length-j-1]
				if len(small) == 0 {
					for _, bi := range big {
						mid := Clone(hash[j][1][0])
						mid.Right = Clone(bi)
						result = append(result, mid)
					}
				} else if len(big) == 0 {
					for _, sm := range small {
						mid := Clone(hash[j][1][0])
						mid.Left = Clone(sm)
						result = append(result, mid)
					}
				} else {

					for _, sm := range small {
						for _, bi := range big {
							mid := Clone(hash[j][1][0])
							mid.Right = Clone(bi)
							mid.Left = Clone(sm)
							result = append(result, mid)
						}
					}
				}

			}
			fmt.Println(len(result))
			hash[i][length] = result
		}
	}
	//	fmt.Println(hash)
	return hash[1][n]
}

func main() {
	printNode(generateTrees(3)...)
}
