package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Priority Queue(Heap Sort)
func getOrder(tasks [][]int) []int {
	n := len(tasks)
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		a, b := idxs[i], idxs[j]
		if tasks[a][0] != tasks[b][0] {
			return tasks[a][0] < tasks[b][0]
		}
		return a < b
	})
	t := 0
	idx := 0
	out := make([]int, n)
	pq := PriorityQueue{tasks: tasks}
	for len(idxs) > 0 || pq.Len() > 0 {
		for len(idxs) > 0 && tasks[idxs[0]][0] <= t {
			i := idxs[0]
			idxs = idxs[1:]
			heap.Push(&pq, i)
		}
		if pq.Len() == 0 && len(idxs) > 0 {
			t = tasks[idxs[0]][0]
		}
		if pq.Len() > 0 {
			i := heap.Pop(&pq).(int)
			out[idx] = i
			idx++
			t += tasks[i][1]
		}
	}
	return out
}

type PriorityQueue struct {
	tasks [][]int
	idxs  []int
}

func (pq PriorityQueue) Len() int { return len(pq.idxs) }
func (pq PriorityQueue) Less(i, j int) bool {
	a, b := pq.idxs[i], pq.idxs[j]
	if pq.tasks[a][1] != pq.tasks[b][1] {
		return pq.tasks[a][1] < pq.tasks[b][1]
	}
	return a < b
}
func (pq PriorityQueue) Swap(i, j int)       { pq.idxs[i], pq.idxs[j] = pq.idxs[j], pq.idxs[i] }
func (pq *PriorityQueue) Push(x interface{}) { pq.idxs = append(pq.idxs, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := pq.idxs[top]
	pq.idxs = pq.idxs[:top]
	return out
}
