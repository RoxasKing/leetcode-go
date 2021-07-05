package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Prefix Sum + Binary Search + Two Pointers
func rangeSum(nums []int, n int, left int, right int) int {
	pSum := make([]int, n+1)
	ppSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
		ppSum[i+1] = ppSum[i] + pSum[i+1]
	}
	mod := int(1e9 + 7)
	return (theKthSum(pSum, ppSum, n, right) - theKthSum(pSum, ppSum, n, left-1) + mod) % mod
}

func theKthSum(pSum, ppSum []int, n, k int) int {
	num := bSearchTheKSum(pSum, n, k)
	sum, cnt := 0, 0
	for i, j := 0, 0; i <= n; i++ {
		for j+1 <= n && pSum[j+1]-pSum[i] < num {
			j++
		}
		sum += ppSum[j] - ppSum[i] - pSum[i]*(j-i)
		cnt += j - i
	}
	return sum + (k-cnt)*num
}

func bSearchTheKSum(pSum []int, n, k int) int {
	l, r := 0, pSum[n]
	for l < r {
		sum := (l + r) >> 1
		cnt := cntNoBiggerSum(pSum, n, sum)
		if cnt < k {
			l = sum + 1
		} else {
			r = sum
		}
	}
	return l
}

func cntNoBiggerSum(pSum []int, n, sum int) int {
	out := 0
	for i, j := 0, 0; i <= n; i++ {
		for j+1 <= n && pSum[j+1]-pSum[i] <= sum {
			j++
		}
		out += j - i
	}
	return out
}

// Prefix Sum + Priority Queue + Merge Sort
func rangeSum2(nums []int, n int, left int, right int) int {
	pq := PriorityQueue{}
	for i := 0; i < n; i++ {
		heap.Push(&pq, &prefixSum{sum: nums[i], idx: i})
	}
	idx := 0
	sum := 0
	for idx < right {
		idx++
		ps := heap.Pop(&pq).(*prefixSum)
		if idx >= left {
			sum += ps.sum
			sum %= 1e9 + 7
		}
		ps.idx++
		if ps.idx == n {
			continue
		}
		ps.sum += nums[ps.idx]
		heap.Push(&pq, ps)
	}
	return sum
}

type prefixSum struct {
	sum int
	idx int
}

type PriorityQueue []*prefixSum

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].sum < pq[j].sum }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*prefixSum)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}

// Prefix Sum
func rangeSum3(nums []int, n int, left int, right int) int {
	pSum := make([]int, n)
	sums := make([]int, 0, n*(n+1)/2)
	for i := 0; i < n; i++ {
		pSum[i] = nums[i]
		if i > 0 {
			pSum[i] += pSum[i-1]
		}
		sums = append(sums, pSum[i])
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sums = append(sums, pSum[j]-pSum[i])
		}
	}
	sort.Ints(sums)
	sum := 0
	for i := left - 1; i < right; i++ {
		sum += sums[i]
		sum %= 1e9 + 7
	}
	return sum
}
