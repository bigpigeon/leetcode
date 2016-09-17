/*
67. Add Binary

Given two binary strings, return their sum (also a binary string).

For example,
a = "11"
b = "1"
Return "100".
*/
package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	carry := byte(0)
	s := []byte{}
	i, j := len(a)-1, len(b)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {

		n := byte(0)
		if a[i] == '1' {
			n++
		}
		if b[j] == '1' {
			n++
		}
		if carry == 1 {
			n++
		}
		s = append(s, n%2+'0')
		carry = n / 2

	}
	for ; i >= 0; i-- {
		n := byte(0)
		if a[i] == '1' {
			n++
		}
		if carry == 1 {
			n++
		}
		s = append(s, n%2+'0')
		carry = n / 2
	}

	for ; j >= 0; j-- {
		n := byte(0)
		if b[j] == '1' {
			n++
		}
		if carry == 1 {
			n++
		}
		s = append(s, n%2+'0')
		carry = n / 2
	}

	if carry == 1 {
		s = append(s, '1')
	}
	//reverse the string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(addBinary("11", "11"))
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("11", "110"))
}
