package main

import (
	"container/heap"
	"sort"
)

/*
  Given a non-empty list of words, return the k most frequent elements.

  Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

  Example 1:
    Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
    Output: ["i", "love"]
    Explanation: "i" and "love" are the two most frequent words.
        Note that "i" comes before "love" due to a lower alphabetical order.

  Example 2:
    Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
    Output: ["the", "is", "sunny", "day"]
    Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
        with the number of occurrence being 4, 3, 2 and 1 respectively.

  Note:
    1. You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
    2. Input words contain only lowercase letters.

  Follow up:
    Try to solve it in O(n log k) time and O(n) extra space.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/top-k-frequent-words
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
