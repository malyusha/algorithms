package main

func twoSum(nums []int, target int) []int {
	var cached = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]

		if prev, ok := cached[diff]; ok {
			return []int{prev, i}
		}

		cached[nums[i]] = i
	}

	return []int{}
}
