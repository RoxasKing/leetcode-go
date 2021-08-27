package main

import "container/heap"

// Priority Queue
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	sum := float64(0)
	pq := PriorityQueue{}
	for _, c := range classes {
		if c[0] == c[1] {
			sum += 1.0
		} else {
			heap.Push(&pq, [2]int{c[0], c[1]})
		}
	}
	for extraStudents > 0 && pq.Len() > 0 {
		e := heap.Pop(&pq).([2]int)
		heap.Push(&pq, [2]int{e[0] + 1, e[1] + 1})
		extraStudents--
	}
	for _, class := range pq {
		sum += float64(class[0]) / float64(class[1])
	}
	return sum / float64(len(classes))
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	inc1 := float64(pq[i][0]+1)/float64(pq[i][1]+1) - float64(pq[i][0])/float64(pq[i][1])
	inc2 := float64(pq[j][0]+1)/float64(pq[j][1]+1) - float64(pq[j][0])/float64(pq[j][1])
	return inc1 > inc2
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([2]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := len(*pq) - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
