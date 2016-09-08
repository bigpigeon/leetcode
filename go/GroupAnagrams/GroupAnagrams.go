/*
49. Group Anagrams

Given an array of strings, group anagrams together.

For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:
+-----------------------+
|[                      |
|  ["ate", "eat","tea"],|
|  ["nat","tan"],       |
|  ["bat"]              |
|]                      |
|                       |
+-----------------------+
Note: All inputs will be in lower-case.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func groupAnagrams(strs []string) [][]string {
	wordCollection := map[[26]byte]int{}
	solution := [][]string{}
	for _, s := range strs {
		w := [26]byte{}
		for i := 0; i < len(s); i++ {
			w[s[i]-'a']++
		}
		if _, ok := wordCollection[w]; !ok {
			wordCollection[w] = len(solution)
			solution = append(solution, []string{s})
		} else {
			index := wordCollection[w]
			solution[index] = append(solution[index], s)
		}
	}
	return solution
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
