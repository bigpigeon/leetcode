/*
127. Word Ladder
Hard

A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:

    Every adjacent pair of words differs by a single letter.
    Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
    sk == endWord

Given two words, beginWord and endWord, and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, or 0 if no such sequence exists.



Example 1:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: One shortest transformation sequence is "hit" -> "hot" -> "dot" -> "dog" -> cog", which is 5 words long.

Example 2:

Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.



Constraints:

    1 <= beginWord.length <= 10
    endWord.length == beginWord.length
    1 <= wordList.length <= 5000
    wordList[i].length == beginWord.length
    beginWord, endWord, and wordList[i] consist of lowercase English letters.
    beginWord != endWord
    All the words in wordList are unique.



*/
package main

import "fmt"

// how match character equal
func wordDiff(word1 string, word2 string) int {
	diff := 0
	for i := range word1 {
		if word1[i] != word2[i] {
			diff++
		}
	}
	return diff
}

func findMatch(closestIdxList [][]int, nextList []int, usedList map[int]struct{}) [][]int {
	if len(nextList) == 0 {
		return nil
	}
	var subNextList []int
	nextListInverted := map[int][]int{}
	var rst [][]int
	for _, next := range nextList {
		closestIdx := closestIdxList[next]
		for _, idx := range closestIdx {
			if idx == -1 {
				rst = append(rst, []int{next})
				break
			}
		}
	}

	if len(rst) != 0 {
		return rst
	}
	for _, next := range nextList {
		if _, ok := usedList[next]; ok {
			continue
		}
		closestIdx := closestIdxList[next]
		for _, subNext := range closestIdx {
			if nextListInverted[subNext] == nil {
				subNextList = append(subNextList, subNext)
			}
			nextListInverted[subNext] = append(nextListInverted[subNext], next)
		}
		usedList[next] = struct{}{}
	}

	subMatch := findMatch(closestIdxList, subNextList, usedList)
	match := [][]int{}
	for _, v := range subMatch {
		last := v[len(v)-1]
		nextList := nextListInverted[last]
		for _, next := range nextList {
			wordLine := append([]int{}, v...)
			wordLine = append(wordLine, next)
			match = append(match, wordLine)
		}
	}
	return match
}

func FindMatch(closestIdxList [][]int, endPos int) [][]int {
	match := findMatch(closestIdxList, []int{endPos}, map[int]struct{}{})
	return match
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	closestIdxList := make([][]int, len(wordList))
	endWordIdx := -1
	for i, word := range wordList {
		if wordDiff(word, beginWord) <= 1 {
			closestIdxList[i] = append(closestIdxList[i], -1)
		}
		for j, preWord := range wordList[:i] {
			if wordDiff(word, preWord) <= 1 {
				closestIdxList[i] = append(closestIdxList[i], j)
			}
		}
		for j, afterWord := range wordList[i+1:] {
			if wordDiff(word, afterWord) <= 1 {
				closestIdxList[i] = append(closestIdxList[i], j+i+1)
			}
		}
		if word == endWord {
			endWordIdx = i
		}
	}
	if endWordIdx == -1 {
		return nil
	}

	matchIdxList := FindMatch(closestIdxList, endWordIdx)
	var closestWordList [][]string

	for _, _lst := range matchIdxList {
		wordLine := []string{beginWord}
		for i := 0; i < len(_lst); i++ {
			wordLine = append(wordLine, wordList[_lst[i]])
		}
		closestWordList = append(closestWordList, wordLine)
	}

	return closestWordList
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	rst := findLadders(beginWord, endWord, wordList)
	if len(rst) != 0 {
		return len(rst[0])
	}
	return 0
}

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
