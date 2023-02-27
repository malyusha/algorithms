package main

func pivotIndex(nums []int) int {
	var sum, leftsum int

	for _, n := range nums {
		sum += n
	}

	for i := 0; i < len(nums); i++ {
		if leftsum == sum-leftsum-nums[i] {
			return i
		}

		leftsum += nums[i]
	}

	return -1
}
