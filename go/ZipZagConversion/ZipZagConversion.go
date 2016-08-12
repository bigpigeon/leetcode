package main

import (
	"fmt"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)
+-------------+
|P   A   H   N|
|A P L S I I G|
|Y   I   R    |
+-------------+
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
+-----------------------------------------+
|  string convert(string text, int nRows);|
+-----------------------------------------+
convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

think about rule:
zigzag patton table when number rows = 2
+---------------+
|0 2 4 6 8 10 12|
|1 3 5 7 9 11 13|
+---------------+


zigzag patton table when number rows = 3
+-----------------+
|0   4   8     12 |
|1 3 5 7 9  11 13 |
|2   6   10       |
+-----------------+
zigzag patton table when number rows = 5
+--------------+
|0	     8     |
|1     7 9     |
|2   6   10    |
|3 5     11 13 |
|4       12    |
+--------------+

*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	PrintCache := make([][]byte, numRows)
	for r, i := 0, 0; i != len(s); i++ {
		PrintCache[r] = append(PrintCache[r], s[i])
		if i/(numRows-1)%2 == 0 {
			r++
		} else {
			r--
		}
	}
	// print cache
	//	for _, line := range PrintCache {
	//		fmt.Println(string(line))
	//	}
	// compression to one line
	toLine := make([]byte, 0, len(s))
	for _, bs := range PrintCache {
		toLine = append(toLine, bs...)
	}
	return string(toLine)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 1))
	fmt.Println(convert("ABC", 2))
}
