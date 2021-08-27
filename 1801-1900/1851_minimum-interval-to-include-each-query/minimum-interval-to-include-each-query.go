package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Priority Queue
func minInterval(intervals [][]int, queries []int) []int {
	m, n := len(intervals), len(queries)

	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool { return queries[idxs[i]] < queries[idxs[j]] })
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	idx := 0
	pq := MinHeap{}
	out := make([]int, n)

	for _, i := range idxs {
		query := queries[i]
		for ; idx < m && intervals[idx][0] <= query; idx++ {
			heap.Push(&pq, &interval{length: intervals[idx][1] - intervals[idx][0] + 1, right: intervals[idx][1]})
		}

		for pq.Len() > 0 && pq[0].right < query {
			heap.Pop(&pq)
		}

		if pq.Len() > 0 {
			out[i] = pq[0].length
		} else {
			out[i] = -1
		}
	}

	return out
}

type interval struct {
	length, right int
}

type MinHeap []*interval

func (pq MinHeap) Len() int            { return len(pq) }
func (pq MinHeap) Less(i, j int) bool  { return pq[i].length < pq[j].length }
func (pq MinHeap) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MinHeap) Push(x interface{}) { *pq = append(*pq, x.(*interval)) }
func (pq *MinHeap) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
