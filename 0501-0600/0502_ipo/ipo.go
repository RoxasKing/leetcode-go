package main

import "container/heap"

// Tags:
// Greedy
// Priority Queue

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	out := w
	maxh := MaxHeap{}
	minh := MinHeap{}
	for i := range profits {
		p := project{profits: profits[i], capital: capital[i]}
		if capital[i] > w {
			heap.Push(&minh, p)
		} else {
			heap.Push(&maxh, p)
		}
	}
	for ; k > 0; k-- {
		for minh.Len() > 0 && minh[0].capital <= out {
			heap.Push(&maxh, heap.Pop(&minh))
		}
		if maxh.Len() == 0 {
			break
		}
		out += heap.Pop(&maxh).(project).profits
	}
	return out
}

type project struct {
	profits, capital int
}

type MaxHeap []project

func (pq MaxHeap) Len() int            { return len(pq) }
func (pq MaxHeap) Less(i, j int) bool  { return pq[i].profits > pq[j].profits }
func (pq MaxHeap) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MaxHeap) Push(x interface{}) { *pq = append(*pq, x.(project)) }
func (pq *MaxHeap) Pop() interface{} {
	i := pq.Len() - 1
	out := (*pq)[i]
	*pq = (*pq)[:i]
	return out
}

type MinHeap []project

func (pq MinHeap) Len() int            { return len(pq) }
func (pq MinHeap) Less(i, j int) bool  { return pq[i].capital < pq[j].capital }
func (pq MinHeap) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *MinHeap) Push(x interface{}) { *pq = append(*pq, x.(project)) }
func (pq *MinHeap) Pop() interface{} {
	i := pq.Len() - 1
	out := (*pq)[i]
	*pq = (*pq)[:i]
	return out
}
