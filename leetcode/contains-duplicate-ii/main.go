package main

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == len(nums)-1 && nums[0] == nums[len(nums)-1] {
		return true
	}
	set := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if i > k {
			delete(set, nums[i-k-1])
		}

		if _, ok := set[nums[i]]; ok {
			return true
		}

		set[nums[i]] = struct{}{}
	}

	return false
}
