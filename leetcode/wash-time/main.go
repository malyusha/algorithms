package main

import (
	"container/heap"
	"fmt"
)

// min heap
type washes []int

func (w washes) Len() int {
	return len(w)
}

func (w washes) Less(i, j int) bool {
	return w[i] < w[j]
}

func (w washes) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w *washes) Push(x any) {
	*w = append(*w, x.(int))
}

func (w *washes) Pop() any {
	list := *w
	l := len(list)
	el := list[l-1]
	*w = list[:l-1]
	return el
}

func main() {
	fmt.Println(washTime([]int{7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 7}, 5))
}

func washTime(q []int, n int) int {
	w := make(washes, 0, n)

	sum := 0
	for _, t := range q {

		if w.Len() == n {
			sum = heap.Pop(&w).(int)
		}
		heap.Push(&w, sum+t)
	}

	max := w[0]
	for _, t := range w[1:] {
		if t > max {
			max = t
		}
	}
	return max
}
