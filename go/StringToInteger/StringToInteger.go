/*
Implement atoi to convert a string to an integer.

Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

Update (2015-02-10):
The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.

spoilers alert... click to show requirements for atoi.

Subscribe to see which companies asked this question

If no valid conversion could be performed, a zero value is returned. If the correct value
is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648)
*/
package main

import (
	"fmt"
)

// full-bit integer
//const MaxUint = ^uint(0)
//const MaxInt = int(MaxUint >> 1)
//const MinInt = -MaxInt - 1

//64-bit integer
//const MinInt = -1 << 63
//const MaxInt = -(MinInt + 1)

// 32-bit integer
const MinInt = -1 << 31
const MaxInt = -(MinInt + 1)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	for i := 0; i != len(str); i++ {
		if str[i] != ' ' {
			str = str[i:]
			break
		}
	}
	sign := 1
	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			sign = -1
		}
		str = str[1:]
	}
	digit := 0
	for _, c := range str {
		if c >= '0' && c <= '9' {
			single := int(c - '0')
			if sign > 0 {
				if digit > (MaxInt-single)/10 {
					return MaxInt
				}

			} else {
				if digit < (MinInt+single)/10 {
					return MinInt
				}
			}
			digit = digit*10 + single*sign
		} else {
			break
		}
	}
	return digit
}

func main() {
	fmt.Printf("max int:%d,min int:%d\n", MaxInt, MinInt)
	fmt.Println(myAtoi("+23"))
	fmt.Println(myAtoi("-100"))
	fmt.Println(myAtoi("32032"))
	fmt.Println(myAtoi("212b2"))
	fmt.Println(myAtoi("   -212b2"))
	fmt.Println(myAtoi("   -999999999999"))
	fmt.Println(myAtoi("   9223372036854775805"))
	fmt.Println(myAtoi("   2147483645"))
	fmt.Println(myAtoi("   -2147483645"))
	fmt.Println(myAtoi("   -9223372036854775807"))
}
