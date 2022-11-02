package main

import (
	"fmt"
	"container/heap"
)

func connectSticks(sticks []int) int {
	h := IntHeap(sticks)
	heap.Init(&h)
	sum := 0
	for h.Len() > 1 {
		v1 := heap.Pop(&h).(int)
		v2 := heap.Pop(&h).(int)
		fmt.Println(v1, v2)
		v := v1 + v2
		sum += v
		heap.Push(&h, v)
	}
	return sum
}

type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(connectSticks([]int{2,4,3,1,8}))
}