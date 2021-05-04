package main

import (
	"container/heap"
)

/*
  Given a string S, check if the letters can be rearranged so that two acters that are adjacent to each other are not the same.

  If possible, output any possible result.  If not possible, return the empty string.

  Example 1:
    Input: S = "aab"
    Output: "aba"

  Example 2:
    Input: S = "aaab"
    Output: ""

  Note:
    S will consist of lowercase letters and have length in range [1, 500].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reorganize-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
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
