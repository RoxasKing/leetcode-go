package main

import (
	"container/heap"
	"sort"
)

// Tags:
// Trie
// Priority Queue

func topKFrequent(words []string, k int) []string {
	t := &Trie{}
	arr := []*Trie{}
	for _, word := range words {
		ret := t.Insert(word)
		if ret != nil {
			arr = append(arr, ret)
		}
	}

	pq := PriorityQueue{}
	for _, t := range arr {
		heap.Push(&pq, &wordFreq{word: t.word, freq: t.freq})
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

type Trie struct {
	child [26]*Trie
	word  string
	freq  int
}

func (t *Trie) Insert(word string) *Trie {
	for i := range word {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}

	out := t
	if t.freq > 0 {
		out = nil
	}

	t.word = word
	t.freq++

	return out
}

type wordFreq struct {
	word string
	freq int
}

type PriorityQueue []*wordFreq

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq != pq[j].freq {
		return pq[i].freq < pq[j].freq
	}
	return pq[i].word > pq[j].word
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*wordFreq)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
