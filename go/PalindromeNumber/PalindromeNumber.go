package main

import (
	"fmt"
)

/*
9. Palindrome Number
Determine whether an integer is a palindrome. Do this without extra space.


Some hints:
Could negative integers be palindromes? (ie, -1)

If you are thinking of converting the integer to string, note the restriction of using extra space.

You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?

There is a more generic way of solving this problem.
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	val_list := []int{}
	for x != 0 {
		val_list = append(val_list, x%10)
		x = x / 10
	}
	fmt.Println(val_list)
	if len(val_list)%2 == 1 {
		mid := len(val_list) / 2
		for i := 1; i < mid+1; i++ {
			if val_list[mid-i] != val_list[mid+i] {
				return false
			}
		}
	} else {
		for left, right := len(val_list)/2-1, len(val_list)/2; left > -1; left, right = left-1, right+1 {
			if val_list[left] != val_list[right] {
				return false
			}
		}
	}
	return true
}

func main() {
	//	fmt.Println(isPalindrome(0))
	//	fmt.Println(isPalindrome(1))
	//	fmt.Println(isPalindrome(-2))
	fmt.Println(isPalindrome(100))
	//	fmt.Println(isPalindrome(12345))
	//	fmt.Println(isPalindrome(123456))
	//	fmt.Println(isPalindrome(1234321))
	//	fmt.Println(isPalindrome(12321))

}
