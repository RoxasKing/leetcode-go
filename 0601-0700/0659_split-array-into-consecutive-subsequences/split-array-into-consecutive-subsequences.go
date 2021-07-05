package main

import "container/heap"

// Priority Queue(Heap Sort) + Hash
func isPossible(nums []int) bool {
	pqs := make(map[int]*priorityQueue)
	for _, num := range nums {
		if pqs[num] == nil {
			pqs[num] = new(priorityQueue)
		}
		if pq := pqs[num-1]; pq != nil {
			prevLen := heap.Pop(pq).(int)
			if pq.Len() == 0 {
				delete(pqs, num-1)
			}
			heap.Push(pqs[num], prevLen+1)
		} else {
			heap.Push(pqs[num], 1)
		}
	}

	for _, pq := range pqs {
		if (*pq)[0] < 3 {
			return false
		}
	}

	return true
}

type priorityQueue []int

func (p priorityQueue) Len() int            { return len(p) }
func (p priorityQueue) Less(i, j int) bool  { return p[i] < p[j] }
func (p priorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *priorityQueue) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *priorityQueue) Pop() interface{} {
	last := len(*p) - 1
	out := (*p)[last]
	*p = (*p)[:last]
	return out
}
