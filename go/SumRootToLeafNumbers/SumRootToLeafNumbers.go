/*
129. Sum Root to Leaf Numbers
Medium

You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

    For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.

Return the total sum of all root-to-leaf numbers.

A leaf node is a node with no children.



Example 1:

Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.

Example 2:

Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.



Constraints:

    The number of nodes in the tree is in the range [1, 1000].
    0 <= Node.val <= 9
    The depth of the tree will not exceed 10.


*/

package main

import (
	"fmt"
	. "github.com/bigpigeon/leetcode/go/datatype"
)

func getLeaf(root *TreeNode, numPrefix int) int {
	if root.Left == nil && root.Right == nil {
		return numPrefix*10 + root.Val
	}
	num := 0
	if root.Left != nil {
		num += getLeaf(root.Left, numPrefix*10+root.Val)
	}
	if root.Right != nil {
		num += getLeaf(root.Right, numPrefix*10+root.Val)
	}
	return num
}

func sumNumbers(root *TreeNode) int {
	return getLeaf(root, 0)
}

func main() {
	fmt.Println(sumNumbers(ListToTreeNode([]int{1, 2, 3})))
}
