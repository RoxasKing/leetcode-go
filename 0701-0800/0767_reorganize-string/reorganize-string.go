package main

import (
	"container/heap"
)

// Tags:
// Priority Queue
func reorganizeString(S string) string {
	n := len(S)

	cnt := [26]int{}
	for i := range S {
		idx := S[i] - 'a'
		if cnt[idx]++; cnt[idx] > (n+1)>>1 {
			return ""
		}
	}

	pq := PriorityQueue{}
	for i := range cnt {
		if cnt[i] == 0 {
			continue
		}
		heap.Push(&pq, &Pair{ch: byte(i) + 'a', cnt: cnt[i]})
	}

	chs := make([]byte, 0, n)
	for pq.Len() > 1 {
		a, b := heap.Pop(&pq).(*Pair), heap.Pop(&pq).(*Pair)
		chs = append(chs, a.ch, b.ch)
		if a.cnt--; a.cnt > 0 {
			heap.Push(&pq, a)
		}
		if b.cnt--; b.cnt > 0 {
			heap.Push(&pq, b)
		}
	}
	if pq.Len() > 0 {
		chs = append(chs, pq.Pop().(*Pair).ch)
	}
	return string(chs)
}

type Pair struct {
	ch  byte
	cnt int
}

type PriorityQueue []*Pair

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].cnt > pq[j].cnt }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Pair)) }
func (pq *PriorityQueue) Pop() interface{} {
	last := len(*pq) - 1
	out := (*pq)[last]
	*pq = (*pq)[:last]
	return out
}
