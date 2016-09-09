/*
50. Pow(x, n)

Implement pow(x, n).
*/
package main

import (
	"fmt"
	"math"
)

func myPowBase(x float64, n int) float64 {
	result := float64(1)
	for n > 0 {
		result *= x
		n--
	}
	for n < 0 {
		result /= x
		n++
	}
	return result
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n > 0 {
		i := 1
		cache := []float64{x}
		for ; i < n/2; i *= 2 {
			x = x * x
			cache = append(cache, x)
		}
		for j := len(cache) - 1; j >= 0; {
			//			fmt.Println("cache", j, (1 << uint(j)), cache, n, i)
			if i+(1<<uint(j)) > n {
				j--
			} else {
				i += (1 << uint(j))
				x *= cache[j]
			}
		}
	} else {
		i := -1
		cache := []float64{x}
		for ; i > n/2; i *= 2 {
			x = x * x
			cache = append(cache, x)
		}
		for j := len(cache) - 1; j >= 0; {
			//			fmt.Println("cache", j, (1 << uint(j)), cache, n, i)
			if i+(-1<<uint(j)) < n {
				j--
			} else {
				i += (-1 << uint(j))
				x *= cache[j]
			}
		}
		x = 1 / x
	}
	return x
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i, myPowBase(2, i), myPow(2, i))
	}
	for i := 0; i < 11; i++ {
		fmt.Println(i, myPowBase(2, -i), myPow(2, -i))
	}
	fmt.Println(math.MinInt32, myPow(2, math.MinInt32))
	fmt.Println(math.MaxInt32, myPow(0.2, math.MaxInt32))
}
