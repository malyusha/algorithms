package main

func search(nums []int, target int) int {
	pivot := findPivotIndex(nums, 0, len(nums)-1)

	leftIx := binarySearch(nums, target, 0, pivot-1)
	if leftIx != -1 {
		return leftIx
	}

	return binarySearch(nums, target, pivot, len(nums)-1)
}

func findPivotIndex(nums []int, low, hi int) int {
	for low <= hi {
		mid := (hi-low)/2 + low
		if nums[mid] > nums[len(nums)-1] {
			low = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return low
}

func binarySearch(nums []int, target int, low, hi int) int {
	for low <= hi {
		mid := (hi-low)/2 + low
		val := nums[mid]
		if val == target {
			return mid
		}

		if val > target {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
