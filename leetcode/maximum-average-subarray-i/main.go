package main

func findMaxAverage(nums []int, k int) float64 {
	res, sum := 0, 0

	for _, num := range nums[:k] {
		sum += num
		res += num
	}

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > res {
			res = sum
		}
	}

	return float64(res) / float64(k)
}
