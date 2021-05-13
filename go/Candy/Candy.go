/*
135. Candy
Hard

There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:

    Each child must have at least one candy.
    Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.



Example 1:

Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:

Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.



Constraints:

    n == ratings.length
    1 <= n <= 2 * 10**4
    0 <= ratings[i] <= 2 * 10**4

*/

package main

import "fmt"

// example:
// ratings [1,2,3,4,5] candy [1,2,3,4,5]
// ratings [1,2,2,3,4] candy [1,2,1,2,3]
// ratings [1,2,2,2,3] candy [1,2,1,1,2]
// ratings [3,2,1,2,3] candy [3,2,1,2,3]
// ratings [1,2,3,2,1] candy [1,2,3,2,1]

// r(x) = ratings[x] value
// c(x) = candy[x] value
// if r[x-1] >= r[x] <= r[x+1]
// c[x] = 1
// if r[x-1] < r[x] > r[x-1]
// c[x] = max(c[x-1]+1, c[x+1]+1)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	candySum := len(ratings)
	downhill := 0
	uphill := 0
	// cache previous uphill value
	// e.g ratings [1,2,4,3,2,1], c(3) = max(c(2),c(4))+1
	// because in downhill c(x) = c(x+1)+1
	// so I need to iterator allover the ratings to get downhill c(4) value
	preUpHill := 0

	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			if downhill != 0 {
				// downhill is over, to calculate preUphill value
				candySum += max(preUpHill, downhill)
				downhill = 0
				preUpHill = 0
			}
			candySum += uphill
			uphill++
		} else if ratings[i+1] < ratings[i] {
			if uphill != 0 {
				preUpHill = uphill
				uphill = 0
			}
			candySum += downhill
			downhill++
		} else { // r[i] == r[i+1], reset all value and calculate previous function value

			candySum += max(preUpHill, downhill)
			candySum += uphill
			preUpHill = 0
			uphill = 0
			downhill = 0
		}
	}

	candySum += max(preUpHill, downhill)
	candySum += uphill

	return candySum
}

func main() {
	//fmt.Println(candy([]int{1, 0, 2}))
	fmt.Println(candy([]int{1, 2, 2})) // [1,2,1] 4
	//fmt.Println(candy([]int{1, 2, 1}))
	//fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	//
	//fmt.Println(candy([]int{1, 5, 4, 3, 2, 1}))
	//fmt.Println(candy([]int{1, 2, 2, 3, 4}))
	fmt.Println(candy([]int{1, 2, 3, 2, 1}))          // [1,2,3,2,1] 9
	fmt.Println(candy([]int{1, 2, 3, 4, 5}))          // [1,2,3,4,5] 15
	fmt.Println(candy([]int{1, 3, 2, 2, 1}))          // [1,2,1,2,1] 7
	fmt.Println(candy([]int{1, 4, 3, 3, 2, 1}))       //[1,2,1,3,2,1] 10
	fmt.Println(candy([]int{1, 2, 3, 4, 3, 2, 1}))    //[1,2,3,4,3,2,1] 16
	fmt.Println(candy([]int{1, 2, 3, 5, 4, 3, 2, 1})) //[1,2,3,5,4,3,2,1] 21
}
