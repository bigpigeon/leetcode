/*
29. Divide Two Integers

Divide two integers without using multiplication, division and mod operator.

If it is overflow, return MAX_INT.

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

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func divide(dividend int, divisor int) int {

	dvd_sign := btoi(dividend > 0)*2 - 1
	dvd := uint(dvd_sign * dividend)
	dvs_sign := btoi(divisor > 0)*2 - 1
	dvs := uint(dvs_sign * divisor)

	multi_dvs := []uint{}
	for d := dvs; d <= dvd; d = d << 1 {
		multi_dvs = append(multi_dvs, d)
	}
	var result uint
	for i := len(multi_dvs) - 1; i >= 0; {
		if dvd >= multi_dvs[i] {
			dvd -= multi_dvs[i]
			result += (1 << uint(i))
		} else {
			i--
		}
	}
	if result > MaxInt {
		if dvd_sign*dvs_sign > 0 {
			return MaxInt
		}
		return MinInt
	}
	return int(result) * dvd_sign * dvs_sign
}

func main() {
	fmt.Println(divide(8, 2))
	fmt.Println(divide(-8, 2))
	fmt.Println(divide(0, 2))
	fmt.Println(divide(100, 2))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(10000000, -78))
	fmt.Println(divide(-2147483648, 1))
}
