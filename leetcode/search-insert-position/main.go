package main

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + ((high - low) / 2)
		val := nums[mid]

		if val == target {
			return mid
		}

		if target > val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
