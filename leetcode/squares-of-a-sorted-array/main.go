package main

// Using two pointers solution. Creating new slice and cursor to track current inserting position.
// Space complexity O(n)
// Time complexity - O(n)
//
func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))

	start, end := 0, len(nums)-1
	for cur := end; cur >= 0; cur-- {
		ssq := nums[start] * nums[start]
		esq := nums[end] * nums[end]

		if ssq > esq {
			out[cur] = ssq
			start++
		} else {
			out[cur] = esq
			end--
		}
	}

	return out
}
