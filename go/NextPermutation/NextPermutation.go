/*
31. Next Permutation

Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
package main

import (
	"fmt"
)

func searchReverse(nums []int, target int) int {
	s, e := 0, len(nums)
	for s < e {
		mid := (e-s)/2 + s
		if nums[mid] <= target {
			e = mid
		} else {
			s = mid + 1
		}
	}
	return s - 1
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	reverse := nums
	for ; i >= 0; i-- {
		//swap i to just larger positive

		if nums[i] < nums[i+1] {
			//			fmt.Println(nums[i], nums[i+1])
			reverse = nums[i+1:]
			p := searchReverse(reverse, nums[i])
			nums[i], reverse[p] = reverse[p], nums[i]
			break
		}
	}
	for s, e := 0, len(reverse)-1; s < e; s, e = s+1, e-1 {
		reverse[s], reverse[e] = reverse[e], reverse[s]
	}

}

func main() {
	//	fmt.Println(searchReverse([]int{7, 3, 5, 1}, 4))
	//	fmt.Println(searchReverse([]int{}, 4))
	//	fmt.Println(searchReverse([]int{8}, 4))
	//	fmt.Println(searchReverse([]int{8, 5, 4, 3, 1}, 4))
	//	fmt.Println(searchReverse([]int{5, 1}, 1))
	array := [][]int{
		{},
		{1, 1, 5},
		{1, 5, 1},
		{1, 2, 3},
		{3, 2, 1},
		{1, 5, 8, 4, 7, 6, 5, 3, 1},
	}
	for _, ele := range array {

		nextPermutation(ele)
		fmt.Println(ele)
	}
}
