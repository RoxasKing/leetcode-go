package main

import (
	"container/heap"
	"math/rand"
)

// Tags:
// Quick Sort
func kClosest(points [][]int, K int) [][]int {
	quickSort(points, 0, len(points)-1, K)
	return points[:K]
}

func quickSort(points [][]int, l, r int, k int) {
	if l == r {
		return
	}
	pivotIndex := l + rand.Intn(r+1-l)
	points[pivotIndex], points[r] = points[r], points[pivotIndex]
	index := l
	for i := l; i < r; i++ {
		if isCloser(points[i], points[r]) {
			points[i], points[index] = points[index], points[i]
			index++
		}
	}
	points[index], points[r] = points[r], points[index]
	if index < k-1 {
		quickSort(points, index+1, r, k)
	} else if index > k-1 {
		quickSort(points, l, index-1, k)
	}
}

func isCloser(a, b []int) bool { return a[0]*a[0]+a[1]*a[1] < b[0]*b[0]+b[1]*b[1] }

// Priority Queue(Heap Sort)
func kClosest2(points [][]int, K int) [][]int {
	n := len(points)
	pq := make(PriorityQueue, n)
	for i := range pq {
		pq[i] = point{
			val:  points[i][0]*points[i][0] + points[i][1]*points[i][1],
			pair: points[i],
		}
	}
	heap.Init(&pq)
	out := make([][]int, 0, K)
	for len(out) < K {
		out = append(out, heap.Pop(&pq).(point).pair)
	}
	return out
}

type point struct {
	val  int
	pair []int
}
type PriorityQueue []point

func (p PriorityQueue) Len() int            { return len(p) }
func (p PriorityQueue) Less(i, j int) bool  { return p[i].val < p[j].val }
func (p PriorityQueue) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *PriorityQueue) Push(x interface{}) { *p = append(*p, x.(point)) }
func (p *PriorityQueue) Pop() interface{} {
	last := len(*p) - 1
	out := (*p)[last]
	*p = (*p)[:last]
	return out
}
