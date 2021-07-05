package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Greedy Algorithm + Binary Search + Priority Queue
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	out := 0
	pq := &PriorityQueue{}
	mark := make(map[int]bool)
	nums := []int{}

	for i := range nums1 {
		diff := Abs(nums1[i] - nums2[i])
		out = (out + diff) % (1e9 + 7)
		heap.Push(pq, [2]int{diff, i})
		if !mark[nums1[i]] {
			mark[nums1[i]] = true
			nums = append(nums, nums1[i])
		}
	}

	if out == 0 {
		return 0
	}

	sort.Ints(nums)

	max := 0
	for pq.Len() > 0 {
		e := heap.Pop(pq).([2]int)
		diff, j := e[0], e[1]
		if Abs(diff) <= max {
			break
		}

		i := sort.Search(len(nums), func(i int) bool { return nums[i] >= nums2[j] })

		if i >= 0 && i < len(nums) {
			max = Max(max, Abs(diff)-Abs(nums[i]-nums2[j]))
		}
		if i-1 >= 0 {
			max = Max(max, Abs(diff)-Abs(nums[i-1]-nums2[j]))
		}
		if i+1 < len(nums) {
			max = Max(max, Abs(diff)-Abs(nums[i+1]-nums2[j]))
		}
	}

	return out - max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i][0] > pq[j][0] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([2]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
