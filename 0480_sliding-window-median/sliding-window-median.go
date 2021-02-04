package main

import (
	"container/heap"
)

/*
  Median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value. So the median is the mean of the two middle value.

  Examples:
  [2,3,4] , the median is 3

  [2,3], the median is (2 + 3) / 2 = 2.5

  Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Your job is to output the median array for each window in the original array.

  For example,
  Given nums = [1,3,-1,-3,5,3,6,7], and k = 3.

  Window position                Median
  ---------------               -----
  [1  3  -1] -3  5  3  6  7       1
   1 [3  -1  -3] 5  3  6  7       -1
   1  3 [-1  -3  5] 3  6  7       -1
   1  3  -1 [-3  5  3] 6  7       3
   1  3  -1  -3 [5  3  6] 7       5
   1  3  -1  -3  5 [3  6  7]      6
  Therefore, return the median sliding window as [1,-1,-1,3,5,6].

  Note:
  You may assume k is always valid, ie: k is always smaller than input array's size for non-empty array.
  Answers within 10^-5 of the actual value will be accepted as correct.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sliding-window-median
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue + Sliding Window
func medianSlidingWindow(nums []int, k int) []float64 {
	n := len(nums)
	out := make([]float64, 0, n+1-k)
	midh := NewMidHeap()
	for i := range nums {
		midh.Push(nums[i])
		if i > k-1 {
			midh.Erase(nums[i-k])
		}
		if i >= k-1 {
			out = append(out, midh.GetMid())
		}
	}
	return out
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}

type MidHeap struct {
	mark map[int]int // count the number of times the number needs to be deleted
	maxh MaxHeap
	minh MinHeap
	maxd int // count the number of times needs to be deleted in max heap
	mind int // count the number of times needs to be deleted in min heap
}

func NewMidHeap() *MidHeap {
	return &MidHeap{
		mark: make(map[int]int),
		maxh: MaxHeap{},
		minh: MinHeap{},
	}
}

func (h *MidHeap) Push(x int) {
	if h.maxh.Len() == 0 || x <= h.maxh[0] {
		heap.Push(&h.maxh, x)
	} else {
		heap.Push(&h.minh, x)
	}
	h.balance()
}

func (h *MidHeap) Erase(x int) {
	h.mark[x]++
	if x <= h.maxh[0] {
		h.maxd++
	} else {
		h.mind++
	}
	h.balance()
}

func (h *MidHeap) balance() {
	for {
		for h.maxh.Len() > 0 && h.mark[h.maxh[0]] > 0 {
			h.mark[h.maxh[0]]--
			h.maxd--
			heap.Pop(&h.maxh)
		}
		for h.minh.Len() > 0 && h.mark[h.minh[0]] > 0 {
			h.mark[h.minh[0]]--
			h.mind--
			heap.Pop(&h.minh)
		}
		if h.maxh.Len()-h.maxd < h.minh.Len()-h.mind {
			heap.Push(&h.maxh, heap.Pop(&h.minh))
		} else if h.maxh.Len()-h.maxd > h.minh.Len()-h.mind+1 {
			heap.Push(&h.minh, heap.Pop(&h.maxh))
		} else {
			break
		}
	}
}

func (h *MidHeap) GetMid() float64 {
	if (h.maxh.Len()-h.maxd+h.minh.Len()-h.mind)&1 == 1 {
		return float64(h.maxh[0])
	}
	return float64(h.maxh[0]+h.minh[0]) / 2.0
}
