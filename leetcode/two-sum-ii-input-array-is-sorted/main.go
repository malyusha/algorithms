package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		val := target - (numbers[l] + numbers[r])
		if val == 0 {
			return []int{l + 1, r + 1}
		}

		if val < 0 {
			r--
		} else {
			l++
		}
	}

	return []int{}
}
