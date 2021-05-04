package main

import (
	"container/heap"
)

/*
  Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

  Example 1:
    Input: nums = [1,1,1,2,2,3], k = 2
    Output: [1,2]

  Example 2:
    Input: nums = [1], k = 1
    Output: [1]

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. k is in the range [1, the number of unique elements in the array].
    3. It is guaranteed that the answer is unique.

  Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/top-k-frequent-elements
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	h := &MinHeap{}
	for k, v := range cnt {
		heap.Push(h, &pair{num: k, cnt: v})
	}

	out := make([]int, 0, k)
	for len(out) < k {
		out = append(out, heap.Pop(h).(*pair).num)
	}
	return out
}

type pair struct {
	num, cnt int
}

type MinHeap []*pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].cnt > h[j].cnt }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*pair)) }
func (h *MinHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
