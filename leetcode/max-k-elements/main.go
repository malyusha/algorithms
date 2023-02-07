package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(findMaxElements([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(findMaxElements([]int{89, 4, 2, 5, 2, 1, 999, 393, 55, 22, 34, 2000}, 2))
}

func findMaxElements(nums []int, k int) []int {
	if k < 1 {
		return []int{}
	}

	if len(nums) <= k {
		return nums
	}
	pivot := pickPivotIndex(len(nums))
	val := nums[pivot]
	var less, equals, max []int
	for _, num := range nums {
		if val == num {
			equals = append(equals, val)
		} else if num > val {
			max = append(max, num)
		} else {
			less = append(less, num)
		}
	}

	if len(max) == k {
		return max
	}

	if len(max) > k {
		return findMaxElements(max, k)
	}

	if len(max)+len(equals) == k {
		return append(equals, max...)
	}

	if len(max)+len(equals) > k {
		return append(max, equals[:k-len(max)]...)
	}

	max = append(max, equals...)
	return append(max, findMaxElements(less, k-len(max))...)
}

func pickPivotIndex(max int) int {
	return rand.Intn(max)
}
