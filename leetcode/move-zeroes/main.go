package main

func moveZeroes(nums []int) {
	i := 0
	for ix := range nums {
		if nums[ix] != 0 {
			nums[i], nums[ix] = nums[ix], nums[i]
			i++
		}
	}
}
