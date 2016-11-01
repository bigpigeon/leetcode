/*
89. Gray Code

The gray code is a binary numeral system where two successive values differ in only one bit.

Given a non-negative integer n representing the total number of bits in the code, print the sequence of gray code. A gray code sequence must begin with 0.

For example, given n = 2, return [0,1,3,2]. Its gray code sequence is:
+------+
|00 - 0|
|01 - 1|
|11 - 3|
|10 - 2|
+------+
Note:
For a given n, a gray code sequence is not uniquely defined.

For example, [0,2,3,1] is also a valid gray code sequence according to the above definition.

For now, the judge is able to judge based on one instance of gray code sequence. Sorry about that.
*/

package main

import (
	"fmt"
)

func grayCode(n int) []int {
	if n <= 0 {
		return []int{0}
	}
	solution := grayCode(n - 1)
	head := 1 << (uint(n) - 1)
	for i := len(solution) - 1; i >= 0; i-- {
		solution = append(solution, head+solution[i])
	}
	return solution
}

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(grayCode(i))
	}
}
