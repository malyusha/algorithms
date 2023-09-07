package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	for _, interval := range intervals {
		lastIx := len(merged) - 1
		if len(merged) == 0 || merged[lastIx][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			last := merged[lastIx]

			last[1] = max(last[1], interval[1])
		}
	}

	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
