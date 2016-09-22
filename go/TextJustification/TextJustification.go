/*
68. Text Justification

Given an array of words and a length L, format the text such that each line has exactly L characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly L characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

For example,
words: ["This", "is", "an", "example", "of", "text", "justification."]
L: 16.

Return the formatted lines as:
+----------------------+
|[                     |
|   "This    is    an",|
|   "example  of text",|
|   "justification.  " |
|]                     |
+----------------------+
Note: Each word is guaranteed not to exceed L in length.
*/
package main

import (
	"fmt"
	"strings"
)

func PrintMultiple(lines []string) {
	for i, l := range lines {
		fmt.Println(i, ": ", len(l), l)
	}
	fmt.Println()
}

func fillLine(words []string, maxWidth int) string {
	words_len := 0
	for _, w := range words {
		words_len += len(w)
	}
	switch len(words) {
	case 1:
		return words[0] + strings.Repeat(" ", maxWidth-words_len)
	default:
		space := (maxWidth - words_len) / (len(words) - 1)
		extraSpace := (maxWidth - words_len) % (len(words) - 1)
		l := words[0]
		i := 1
		for ; i <= extraSpace; i++ {
			l += strings.Repeat(" ", space+1) + words[i]
		}
		for ; i < len(words); i++ {
			l += strings.Repeat(" ", space) + words[i]
		}
		return l
	}
}

func fullJustify(words []string, maxWidth int) []string {
	lines := []string{}
	curr_len := 0
	cache := []string{}
	for _, w := range words {
		//calculative the space was enougth:words length+words number  < maxWidth
		if curr_len+len(w)+len(cache) <= maxWidth {
			cache = append(cache, w)
			curr_len += len(w)
		} else {
			lines = append(lines, fillLine(cache, maxWidth))
			cache = []string{w}
			curr_len = len(w)
		}
	}
	if len(cache) > 0 {
		lines = append(lines, strings.Join(cache, " ")+strings.Repeat(" ", maxWidth-curr_len-len(cache)+1))
	}
	return lines
}

func main() {
	PrintMultiple(fullJustify([]string{"this", "is", "animal"}, 16))
	//	PrintMultiple(fullJustify([]string{""}, 0))
	PrintMultiple(fullJustify([]string{"What", "must", "be", "shall", "be."}, 12))
	PrintMultiple(fullJustify([]string{"My", "momma", "always", "said,", "\"Life", "was", "like", "a", "box", "of", "chocolates.", "You", "never", "know", "what", "you're", "gonna", "get."}, 20))
}
