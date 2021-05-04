package main

import "container/heap"

/*
  The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value and the median is the mean of the two middle values.
    1. For example, for arr = [2,3,4], the median is 3.
    2. For example, for arr = [2,3], the median is (2 + 3) / 2 = 2.5.

  Implement the MedianFinder class:
    1. MedianFinder() initializes the MedianFinder object.
    2. void addNum(int num) adds the integer num from the data stream to the data structure.
    3. double findMedian() returns the median of all elements so far. Answers within 10-5 of the actual answer will be accepted.

  Example 1:
    Input
      ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
      [[], [1], [2], [], [3], []]
    Output
      [null, null, null, 1.5, null, 2.0]
    Explanation
      MedianFinder medianFinder = new MedianFinder();
      medianFinder.addNum(1);    // arr = [1]
      medianFinder.addNum(2);    // arr = [1, 2]
      medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
      medianFinder.addNum(3);    // arr[1, 2, 3]
      medianFinder.findMedian(); // return 2.0

  Constraints:
    1. -10^5 <= num <= 10^5
    2. There will be at least one element in the data structure before calling findMedian.
    3. At most 5 * 10^4 calls will be made to addNum and findMedian.

  Follow up:
    1. If all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?
    2. If 99% of all integer numbers from the stream are in the range [0, 100], how would you optimize your solution?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-median-from-data-stream
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)

type MedianFinder struct {
	maxh MaxHeap
	minh MinHeap
	size int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxh: MaxHeap{},
		minh: MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.maxh, num)
	this.size++
	for {
		if this.maxh.Len() > 0 && this.minh.Len() > 0 && this.maxh[0] > this.minh[0] {
			a := heap.Pop(&this.maxh)
			b := heap.Pop(&this.minh)
			heap.Push(&this.maxh, b)
			heap.Push(&this.minh, a)
		} else if this.maxh.Len() > this.minh.Len()+1 {
			heap.Push(&this.minh, heap.Pop(&this.maxh))
		} else {
			break
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.size&1 == 1 {
		return float64(this.maxh[0])
	}
	return float64(this.maxh[0]+this.minh[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
