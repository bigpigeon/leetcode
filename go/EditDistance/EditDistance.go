/*
72. Edit Distance

Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2. (each operation is counted as 1 step.)

You have the following 3 operations permitted on a word:

a) Insert a character
b) Delete a character
c) Replace a character

tag Dynamic Programming
*/
package main

import (
	"fmt"
)

func min(l ...int) int {
	if len(l) == 0 {
		panic("argument was empty")
	}
	m := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] < m {
			m = l[i]
		}
	}
	return m
}

func printMatrix(m [][]int, word1, word2 string) {
	fmt.Print("\t")
	for _, v := range word2 {
		fmt.Print("\t", string(v))
	}
	fmt.Println()
	for i, x := range m {
		if i > 0 {
			fmt.Print(string(word1[i-1]), "\t")
		} else {
			fmt.Print("\t")
		}
		for _, y := range x {
			fmt.Print(y, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func minDistance(word1 string, word2 string) int {
	m := make([][]int, len(word1)+1)
	for i, _ := range m {
		m[i] = make([]int, len(word2)+1)
	}
	for i := 1; i < len(word1)+1; i++ {
		m[i][0] = i
	}
	for j := 1; j < len(word2)+1; j++ {
		m[0][j] = j
	}

	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				m[i][j] = m[i-1][j-1]
			} else {
				m[i][j] = min(m[i-1][j-1], m[i-1][j], m[i][j-1]) + 1
			}
		}
	}
	//	printMatrix(m, word1, word2)
	return m[len(word1)][len(word2)]
}

func main() {
	//	fmt.Println(minDistance("abcddde", "qqbcdddeabet"))
	//	fmt.Println(minDistance("abcddde", ""))
	fmt.Println(minDistance("mart", "karma"))
	fmt.Println(minDistance("ttazbxcvdben", "aabbccddeett"))
	//	fmt.Println(minDistance("", ""))
}
