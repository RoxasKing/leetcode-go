package main

import "container/heap"

// Priority Queue(Heap Sort)
func lastStoneWeight(stones []int) int {
	pq := &PriorityQueue{}
	for _, stone := range stones {
		heap.Push(pq, stone)
	}
	for pq.Len() > 1 {
		a := heap.Pop(pq).(int)
		b := heap.Pop(pq).(int)
		if a != b {
			heap.Push(pq, Abs(a-b))
		}
	}
	if pq.Len() == 0 {
		return 0
	}
	return heap.Pop(pq).(int)
}

type PriorityQueue []int

func (q PriorityQueue) Len() int            { return len(q) }
func (q PriorityQueue) Less(i, j int) bool  { return q[i] > q[j] }
func (q PriorityQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(int)) }
func (q *PriorityQueue) Pop() interface{} {
	last := len(*q) - 1
	out := (*q)[last]
	*q = (*q)[:last]
	return out
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
