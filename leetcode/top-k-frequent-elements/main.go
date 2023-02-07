package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(topKFrequent([]int{4, 3, 2, 6, 4, 7, 2, 1, 1}, 3))
}

func topKFrequent(nums []int, k int) []int {
	return topKFrequentBucket(nums, k)
}

func topKFrequentMap(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	uniqueNums := make([]int, 0, len(cnt))
	for num := range cnt {
		uniqueNums = append(uniqueNums, num)
	}

	sort.Slice(uniqueNums, func(i, j int) bool {
		return cnt[uniqueNums[i]] > cnt[uniqueNums[j]]
	})

	return uniqueNums[:k]
}

func topKFrequentBucket(nums []int, k int) []int {
	freq := map[int]int{}
	buckets := make([][]int, len(nums)+1)

	for i := 0; i < len(buckets); i++ {
		buckets[i] = make([]int, 0)
	}

	for _, num := range nums {
		freq[num]++
	}

	for num, frequency := range freq {
		buckets[frequency] = append(buckets[frequency], num)
	}

	out := make([]int, 0)

	for i := len(buckets) - 1; i > 0; i-- {
		if len(buckets[i]) != 0 {
			out = append(out, buckets[i]...)
		}
	}

	return out[:k]
}
