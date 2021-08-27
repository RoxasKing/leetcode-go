package main

import "container/heap"

// Priority Queue
type KthLargest struct {
	maxh MaxHeap
	minh MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	maxh := MaxHeap{}
	minh := MinHeap(nums)
	heap.Init(&minh)
	for minh.Len() > k-1 {
		heap.Push(&maxh, heap.Pop(&minh))
	}
	return KthLargest{maxh: maxh, minh: minh}
}

func (this *KthLargest) Add(val int) int {
	if this.minh.Len() > 0 && val > this.minh.Top() {
		heap.Push(&this.maxh, heap.Pop(&this.minh))
		heap.Push(&this.minh, val)
	} else {
		heap.Push(&this.maxh, val)
	}
	return this.maxh.Top()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

type MaxHeap []int

func (mh MaxHeap) Len() int            { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool  { return mh[i] > mh[j] }
func (mh MaxHeap) Swap(i, j int)       { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MaxHeap) Push(x interface{}) { *mh = append(*mh, x.(int)) }
func (mh *MaxHeap) Pop() interface{} {
	last := len(*mh) - 1
	out := (*mh)[last]
	*mh = (*mh)[:last]
	return out
}
func (mh MaxHeap) Top() int { return mh[0] }

type MinHeap []int

func (mh MinHeap) Len() int            { return len(mh) }
func (mh MinHeap) Less(i, j int) bool  { return mh[i] < mh[j] }
func (mh MinHeap) Swap(i, j int)       { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(x interface{}) { *mh = append(*mh, x.(int)) }
func (mh *MinHeap) Pop() interface{} {
	last := len(*mh) - 1
	out := (*mh)[last]
	*mh = (*mh)[:last]
	return out
}
func (mh MinHeap) Top() int { return mh[0] }
