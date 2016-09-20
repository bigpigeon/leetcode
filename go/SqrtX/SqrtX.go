/*
69. Sqrt(x)

Implement int sqrt(int x).

Compute and return the square root of x.
*/
package main

import (
	"fmt"
)

func mySqrt(x int) int {
	maxBit := uint(64)
	minBit := uint(0)
	for minBit < maxBit {
		halfBit := (maxBit-minBit)/2 + minBit
		v := 1 << halfBit
		if v == x {
			minBit, maxBit = halfBit, halfBit
			break
		} else if v > x {
			maxBit = halfBit
		} else {
			minBit = halfBit + 1
		}
	}
	max := 1 << (minBit/2 + minBit%2)
	min := max >> 1
	for min < max {
		half := (max-min)/2 + min
		v := half * half
		if v == x {
			min, max = half, half
			break
		} else if v < x {
			min = half + 1
		} else {
			max = half
		}

	}
	v := min * min
	if v > x {
		return min - 1
	}
	return min
}

func main() {
	fmt.Println(mySqrt(0))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(1024))
	fmt.Println(mySqrt(512))
	fmt.Println(mySqrt(513))
}
