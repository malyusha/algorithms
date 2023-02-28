package main

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	out := make([][]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		id1, id2 := nums1[i][0], nums2[j][0]
		val1, val2 := nums1[i][1], nums2[j][1]
		var id, val int
		if id1 < id2 {
			id = id1
			val = val1
			i++
		} else if id1 > id2 {
			id = id2
			val = val2
			j++
		} else {
			id = id1
			val = val1 + val2
			i++
			j++
		}

		out = append(out, []int{id, val})
	}

	left := nums1
	leftIter := i
	if leftIter >= len(nums1) {
		left = nums2
		leftIter = j
	}

	for leftIter < len(left) {
		out = append(out, []int{left[leftIter][0], left[leftIter][1]})
		leftIter++
	}

	return out
}
