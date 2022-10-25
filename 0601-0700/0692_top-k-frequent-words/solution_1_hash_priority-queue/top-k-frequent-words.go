package main

import (
	"container/heap"
	"sort"
)

// Difficulty:
// Medium

// Tags:
// Hash
// Priority Queue

func topKFrequent(words []string, k int) []string {
	freq := map[string]int{}
	for _, w := range words {
		freq[w]++
	}

	pq := PriorityQueue{}
	for w, c := range freq {
		heap.Push(&pq, &WordFreq{word: w, freq: c})
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}

	sort.Slice(pq, func(i, j int) bool {
		if pq[i].freq != pq[j].freq {
			return pq[i].freq > pq[j].freq
		}
		return pq[i].word < pq[j].word
	})

	out := make([]string, 0, k)
	for _, wf := range pq {
		out = append(out, wf.word)
	}
	return out
}

type WordFreq struct {
	word string
	freq int
}

type PriorityQueue []*WordFreq

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq != pq[j].freq {
		return pq[i].freq < pq[j].freq
	}
	return pq[i].word > pq[j].word
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*WordFreq)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
