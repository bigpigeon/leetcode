/*
43. Multiply Strings

Given two numbers represented as strings, return multiplication of the numbers as a string.

Note:
The numbers can be arbitrarily large and are non-negative.
Converting the input string to integer is NOT allowed.
You should NOT use internal library such as BigInteger.
*/
package main

import (
	"fmt"
)

func strPlus(num1 string, num2 string) string {
	sum := ""
	carry := byte(0)
	i1, i2 := len(num1)-1, len(num2)-1
	for ; i1 >= 0 && i2 >= 0; i1, i2 = i1-1, i2-1 {
		digit := (num1[i1] - '0') + (num2[i2] - '0') + carry
		sum = string(digit%10+'0') + sum
		carry = digit / 10
	}
	for ; i1 >= 0; i1-- {
		digit := num1[i1] - '0' + carry
		sum = string(digit%10+'0') + sum
		carry = digit / 10
	}
	for ; i2 >= 0; i2-- {
		digit := num2[i2] - '0' + carry
		sum = string(digit%10+'0') + sum
		carry = digit / 10
	}
	for carry > 0 {
		sum = string(carry%10+'0') + sum
		carry /= 10
	}
	return sum
}

func multiply(num1 string, num2 string) string {
	result := "0"
	for i := len(num1) - 1; i >= 0; i-- {
		carry := byte(0)
		multi := ""
		for j := len(num2) - 1; j >= 0; j-- {
			digit := (num1[i]-'0')*(num2[j]-'0') + carry
			multi = string(digit%10+'0') + multi
			carry = digit / 10
		}
		for carry > 0 {
			multi = string(carry%10+'0') + multi
			carry /= 10
		}
		for k := len(num1) - 1; k > i; k-- {
			multi += "0"
		}
		result = strPlus(result, multi)
	}
	//remove top left zero
	i := 0
	for ; i < len(result)-1; i++ {
		if result[i] != '0' {
			break
		}
	}

	return result[i:]
}

func main() {
	fmt.Println(multiply("12", "32"))
	fmt.Println(multiply("0", "0"))
	fmt.Println(multiply("0", "320"))
	fmt.Println(multiply("1234567891000", "1234567891000"))
}
