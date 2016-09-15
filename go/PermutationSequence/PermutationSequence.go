/*
60. Permutation Sequence

The set [1,2,3,…,n] contains a total of n! unique permutations.

By listing and labeling all of the permutations in order,
We get the following sequence (ie, for n = 3):

"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.

Note: Given n will be between 1 and 9 inclusive.
*/
package main

import (
	"fmt"
)

func getPermutation(n int, k int) string {
	k--
	s := []byte{}
	candidate := make([]int, 0, n)
	nFactorial := 1
	for i := 1; i <= n; i++ {
		candidate = append(candidate, i)
		nFactorial *= i
	}
	for ; n >= 1; n-- {
		nFactorial /= n
		d := k / nFactorial
		s = append(s, byte(candidate[d])+'0')
		sc := candidate[d+1:]
		candidate = candidate[0:d]
		candidate = append(candidate, sc...)
		k = k % nFactorial
	}
	return string(s)
}

func main() {
	fmt.Println(getPermutation(3, 5))
	fmt.Println(getPermutation(2, 1))
}
