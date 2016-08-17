/*
12. Integer to Roman
Given an integer, convert it to a roman numeral.

Input is guaranteed to be within the range from 1 to 3999.
*/
package main

import (
	"fmt"
)

//type RomanNum struct {
//	//thousand bit
//	M int
//	//hundred bit
//	D int
//	C int
//	//ten bit
//	L int
//	X int
//	//unit bit
//	V int
//	I int
//}

func RomanFormat(ten_symbol, five_symbol, unit_symbal byte, num int) string {
	roman_num := []byte{}
	if num == 9 {
		roman_num = append(roman_num, unit_symbal, ten_symbol)
	} else if num >= 5 {
		roman_num = append(roman_num, five_symbol)
		for i := 0; i < num-5; i++ {
			roman_num = append(roman_num, unit_symbal)
		}
	} else if num == 4 {
		roman_num = append(roman_num, unit_symbal, five_symbol)
	} else {
		for i := 0; i < num; i++ {
			roman_num = append(roman_num, unit_symbal)
		}
	}
	return string(roman_num)
}

func intToRoman(num int) string {
	roman_num := ""
	for i := 0; i < num/1000; i++ {
		roman_num += "M"
	}
	num %= 1000

	hundred_num := num / 100
	roman_num += RomanFormat('M', 'D', 'C', hundred_num)
	num %= 100

	ten_num := num / 10
	roman_num += RomanFormat('C', 'L', 'X', ten_num)
	num %= 10

	roman_num += RomanFormat('X', 'V', 'I', num)
	return roman_num
}

func main() {
	fmt.Println(intToRoman(1999))
	fmt.Println(intToRoman(1000))
	fmt.Println(intToRoman(599))
	fmt.Println(intToRoman(500))
	fmt.Println(intToRoman(50))
	fmt.Println(intToRoman(0))
}
