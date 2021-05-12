/*
138. Copy List with Random Pointer
Medium

A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

    val: an integer representing Node.val
    random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.

Your code will only be given the head of the original linked list.



Example 1:

Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

Example 2:

Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

Example 3:

Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

Example 4:

Input: head = []
Output: []
Explanation: The given linked list is empty (null pointer), so return null.



Constraints:

    0 <= n <= 1000
    -10000 <= Node.val <= 10000
    Node.random is null or is pointing to some node in the linked list.

*/

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	newHead := &Node{Val: head.Val}
	// process random node copy first
	preNodeMap := map[*Node]*Node{head: newHead}
	{
		node := head
		for node != nil {
			if node.Random != nil {

				var randomNode *Node
				if existed := preNodeMap[node.Random]; existed != nil {
					randomNode = existed
				} else {
					randomNode = &Node{Val: node.Random.Val}
					preNodeMap[node.Random] = randomNode
				}
				var hasRandom *Node
				if existed := preNodeMap[node]; existed != nil {
					hasRandom = existed
				} else {
					hasRandom = &Node{Val: node.Val}
				}
				hasRandom.Random = randomNode
				preNodeMap[node] = hasRandom
			}
			node = node.Next
		}
	}

	nextNode := head.Next
	newNode := newHead

	for nextNode != nil {
		if existed := preNodeMap[nextNode]; existed != nil {
			newNode.Next = existed
		} else {
			newNode.Next = &Node{Val: nextNode.Val}
		}
		nextNode = nextNode.Next
		newNode = newNode.Next
	}
	return newHead
}

func printNode(node *Node) {

	for node != nil {
		fmt.Printf("%d %p %p \n", node.Val, node, node.Random)
		node = node.Next
	}
}

func prepareNode1() *Node {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n3 := &Node{Val: 3}
	n4 := &Node{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n3.Random = n1
	n2.Random = n4
	return n1
}

func makeLink(nodes ...*Node) {
	if len(nodes) < 2 {
		return
	}
	pre := nodes[0]
	for _, n := range nodes[1:] {
		pre.Next = n
		pre = n
	}
}

func prepareNode2() *Node {
	n1 := &Node{Val: 7}
	n2 := &Node{Val: 13}
	n3 := &Node{Val: 11}
	n4 := &Node{Val: 10}
	n5 := &Node{Val: 1}
	makeLink(n1, n2, n3, n4, n5)
	n2.Random = n1
	n3.Random = n5
	n4.Random = n3
	n5.Random = n1
	return n1
}

func main() {
	//{
	//	n1 := prepareNode1()
	//	printNode(n1)
	//	copyN1 := copyRandomList(n1)
	//	fmt.Println("---")
	//	printNode(copyN1)
	//}
	fmt.Println("---")
	{
		n1 := prepareNode2()
		printNode(n1)
		copyN1 := copyRandomList(n1)
		fmt.Println("---")
		printNode(copyN1)
	}
}
