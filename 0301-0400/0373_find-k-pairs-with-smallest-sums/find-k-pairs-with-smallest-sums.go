package main

import (
	"container/heap"
	"sort"
)

/*
  You are given two integer arrays nums1 and nums2 sorted in ascending order and an integer k.

  Define a pair (u, v) which consists of one element from the first array and one element from the second array.

  Return the k pairs (u1, v1), (u2, v2), ..., (uk, vk) with the smallest sums.

  Example 1:
    Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
    Output: [[1,2],[1,4],[1,6]]
    Explanation: The first 3 pairs are returned from the sequence: [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

  Example 2:
    Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
    Output: [[1,1],[1,1]]
    Explanation: The first 2 pairs are returned from the sequence: [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

  Example 3:
    Input: nums1 = [1,2], nums2 = [3], k = 3
    Output: [[1,3],[2,3]]
    Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]

  Constraints:
    1. 1 <= nums1.length, nums2.length <= 10^4
    2. -10^9 <= nums1[i], nums2[i] <= 10^9
    3. nums1 and nums2 both are sorted in ascending order.
    4. 1 <= k <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := MaxHeap{}
	for _, u := range nums1 {
		for _, v := range nums2 {
			if h.Len() == k && h[0][0]+h[0][1] <= u+v {
				break
			}
			if h.Len() == k && h[0][0]+h[0][1] > u+v {
				heap.Pop(&h)
			}
			heap.Push(&h, []int{u, v})
		}
	}
	sort.Slice(h, func(i, j int) bool {
		if h[i][0]+h[i][1] != h[j][0]+h[j][1] {
			return h[i][0]+h[i][1] < h[j][0]+h[j][1]
		}
		if h[i][0] != h[j][0] {
			return h[i][0] < h[j][0]
		}
		return h[i][1] < h[j][1]
	})
	return h
}

type MaxHeap [][]int // [u, v] pair

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i][0]+h[i][1] != h[j][0]+h[j][1] {
		return h[i][0]+h[i][1] > h[j][0]+h[j][1]
	}
	if h[i][0] != h[j][0] {
		return h[i][0] > h[j][0]
	}
	return h[i][1] > h[j][1]
}
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *MaxHeap) Pop() interface{} {
	top := h.Len() - 1
	out := (*h)[top]
	*h = (*h)[:top]
	return out
}
