/*
96. Unique Binary Search Trees

Given n, how many structurally unique BST's (binary search trees) that store values 1...n?

For example,
Given n = 3, there are a total of 5 unique BST's.

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

func numTrees(n int) int {
	// hash[start][length] start mean start num, length means end -start
	hash := map[int]map[int]int{}
	for i := 1; i <= n; i++ {
		hash[i] = map[int]int{}
		hash[i][1] = 1
	}
	// full
	hash[n+1] = map[int]int{}

	for length := 2; length <= n; length++ {
		for i := 1; i <= n-length+1; i++ {
			sum := 0
			for j := i; j < i+length; j++ {
				small := hash[i][j-i]
				big := hash[j+1][i+length-j-1]

				// TODO detach
				if small == 0 {
					sum += big
				} else if big == 0 {
					sum += small
				} else {
					sum += big * small
				}
			}
			hash[i][length] = sum
		}
	}
	fmt.Println(hash)
	return hash[1][n]
}

func main() {
	fmt.Println(numTrees(3))
}
