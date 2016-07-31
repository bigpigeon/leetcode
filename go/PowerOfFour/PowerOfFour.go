package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	//equal 0b0101010101010101010101010101010101010101010101010101010101010101
	const mask int = 0x5555555555555555
	//ignore negative
	if num <= 0 {
		return false
	}
	//detect power of two
	//example: bin of eight:1000 bin of seven: 111,
	if num&(num-1) != 0 {
		return false
	}

	//detect power of four
	if (num & mask) != 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPowerOfFour(2))
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(6))
	fmt.Println(isPowerOfFour(8))
	fmt.Println(isPowerOfFour(16))
}
