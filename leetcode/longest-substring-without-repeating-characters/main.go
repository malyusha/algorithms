package main

func lengthOfLongestSubstring(s string) (res int) {
	m := make(map[rune]int)
	curr := 0

	for ix, r := range s {
		if val, ok := m[r]; ok {
			curr = max(curr, val)
		}

		win := ix - curr + 1
		res = max(res, win)

		m[r] = ix + 1
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
