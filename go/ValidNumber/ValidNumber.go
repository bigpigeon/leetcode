/*
65. Valid Number

Validate if a given string is numeric.

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
Note: It is intended for the problem statement to be ambiguous. You should gather all requirements up front before implementing one.

Update (2015-02-10):
The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.
*/
package main

import (
	"fmt"
)

func isSignNumberLen(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		pos := isPositiveLen(s[1:])
		if pos == 0 {
			return 0
		}
		return pos + 1
	}
	return isPositiveLen(s)
}

func isPositiveLen(s string) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			if s[i] == '.' {
				pos := isPositiveIntegerLen(s[i+1:])
				if i == 0 && pos == 0 { // just like '.', '..'
					return 0
				} else {
					return pos + i + 1
				}
			} else {
				return i
			}
		}
	}

	return len(s)
}

func isSignPositiveIntegerLen(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		pos := isPositiveIntegerLen(s[1:])
		if pos == 0 {
			return 0
		}
		return pos + 1
	}
	return isPositiveIntegerLen(s)
}

func isPositiveIntegerLen(s string) int {
	for i, c := range s {
		if c < '0' || c > '9' {
			return i
		}
	}
	return len(s)
}

func isNumber(s string) bool {
	//strip space with head
	h := 0
	for ; h < len(s) && s[h] == ' '; h++ {

	}
	t := len(s)
	for ; t > h && s[t-1] == ' '; t-- {

	}
	s = s[h:t]
	pos := isSignNumberLen(s)
	if pos == 0 {
		return false
	}
	if pos == len(s) {
		return true
	}
	if s[pos] == 'e' {
		new_pos := isSignPositiveIntegerLen(s[pos+1:])
		if new_pos == 0 || new_pos != len(s[pos+1:]) {
			return false
		}
		return true
	}
	return false
}

func main() {
	//	fmt.Println(isFloatLen("23.13"))
	//	fmt.Println(isFloatLen("2323"))
	//	fmt.Println(isFloatLen("-2323"))
	//	fmt.Println(isFloatLen("-2323.23.4141"))
	//	fmt.Println(isFloatLen("-2323.-23.4141"))
	fmt.Println(isNumber(" 0.1 "))
	fmt.Println(isNumber("23"))
	fmt.Println(isNumber("223.2"))
	fmt.Println(isNumber("-223.2"))
	fmt.Println(isNumber("-223.2e23"))
	fmt.Println(isNumber(".1"))

	fmt.Println(isNumber("1 4"))
	fmt.Println(isNumber("-223.2e23-232"))
	fmt.Println(isNumber("-223.2e23.232"))
	fmt.Println(isNumber(""))
	fmt.Println(isNumber("-23e"))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber(".."))
	fmt.Println(isNumber("e"))

}
