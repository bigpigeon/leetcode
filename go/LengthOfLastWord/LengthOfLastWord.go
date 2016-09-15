/*
58. Length of Last Word

Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

For example,
Given s = "Hello World",
return 5.
*/
package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	end := len(s)
	//strip the space of the end
	for ; end > 0; end-- {
		if s[end-1] != ' ' {
			break
		}
	}
	start := end - 1
	for ; start >= 0 && s[start] != ' '; start-- {

	}
	start++
	return end - start
}

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord("helloworld"))
}
