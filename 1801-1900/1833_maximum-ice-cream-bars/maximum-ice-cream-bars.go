package main

import (
	"container/heap"
)

// Tags:
// Priority Queue
func maxIceCream(costs []int, coins int) int {
	pq := PriorityQueue{}
	for _, c := range costs {
		heap.Push(&pq, c)
	}
	out := 0
	for pq.Len() > 0 && coins-pq[0] >= 0 {
		coins -= heap.Pop(&pq).(int)
		out++
	}
	return out
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
