package main

func removeElement(nums []int, val int) int {
	n := 0

	for i := range nums {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}

	return n
}
