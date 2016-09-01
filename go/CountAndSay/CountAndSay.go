/*
38. Count and Say

The count-and-say sequence is the sequence of integers beginning as follows:
1, 11, 21, 1211, 111221, ...

1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
**Given an integer n, generate the nth sequence.**

Note: The sequence of integers will be represented as a string.

sequence verbose:
1, 11, 21, 1211, 111221, 312211, 13112221, ...
*/
package main

import (
	"fmt"
)

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	s := "1"
	for i := 1; i < n; i++ {

		digitCount := []*struct {
			digit byte
			num   int
		}{{s[0], 1}}
		dc := digitCount[len(digitCount)-1]
		for j := 1; j < len(s); j++ {

			if s[j] == dc.digit {
				dc.num++
			} else {
				dc = &struct {
					digit byte
					num   int
				}{s[j], 1}
				digitCount = append(digitCount, dc)
			}
		}
		new_s := ""
		for _, d := range digitCount {
			new_s += string(d.num+'0') + string(d.digit)
		}
		s = new_s
	}
	return s
}

func main() {
	for i := 0; i < 9; i++ {
		fmt.Println(countAndSay(i))
	}
}
