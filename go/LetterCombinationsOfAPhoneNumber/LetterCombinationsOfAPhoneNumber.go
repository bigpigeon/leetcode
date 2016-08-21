/*
17. Letter Combinations of a Phone Number
Given a digit string, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below.
+------------------------+
| 1 nil   2 abc   3 def  |
|                        |
| 4 ghi   5 jkl   6 mno  |
|                        |
| 7 pqrs  8 tuv   9 wxyz |
|                        |
| * +     0 _     ⇧ #    |
+------------------------+

+---------------------------------------------------------------+
|Input:Digit string "23"                                        |
|Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].|
+---------------------------------------------------------------+

Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
*/
package main

import (
	"fmt"
)

var NumberRepresents = map[rune]string{
	'0': " ",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
	'*': "+",
	'⇧': "#",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	runeDigits := []rune(digits)
	d := runeDigits[0]
	if str, exist := NumberRepresents[d]; exist {
		if len(runeDigits) == 1 {
			combinations := make([]string, 0, len(str))
			for _, c := range str {
				combinations = append(combinations, string(c))
			}
			return combinations
		} else {
			tailCombinations := letterCombinations(string(runeDigits[1:]))
			combinations := make([]string, 0, len(tailCombinations)*len(str))
			for _, c := range str {
				for _, v := range tailCombinations {
					combinations = append(combinations, string(append([]rune{c}, []rune(v)...)))
				}
			}
			return combinations
		}
	}
	return []string{}
}

func main() {
	fmt.Println(letterCombinations("2362"))
	fmt.Println(letterCombinations("13"))
	fmt.Println(letterCombinations("*25"))
	fmt.Println(letterCombinations("⇧"))
}
