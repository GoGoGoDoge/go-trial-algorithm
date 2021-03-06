package kth_largest_element_in_a_stream

import "container/heap"

type KthLargest struct {
	heap intHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	heap.Init(&h)

	for len(h) > k {
		heap.Pop(&h)
	}

	return KthLargest{
		heap: h,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.heap, val)

	if len(this.heap) > this.k {
		heap.Pop(&this.heap)
	}

	return this.heap[0]
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	val := x.(int)
	*h = append(*h, val)
}
func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
