package main

import "container/heap"

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

// Hash + Priority Queue
func topKFrequent(words []string, k int) []string {
	wCnt := make(map[string]int)
	for _, w := range words {
		wCnt[w]++
	}
	pq := PriorityQueue{}
	for w, c := range wCnt {
		heap.Push(&pq, &WordFreq{word: w, freq: c})
	}
	out := make([]string, 0, k)
	for k > 0 {
		w := heap.Pop(&pq).(*WordFreq).word
		out = append(out, w)
		k--
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
		return pq[i].freq > pq[j].freq
	}
	return pq[i].word < pq[j].word
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*WordFreq)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}

// Trie + Priority Queue
func topKFrequent2(words []string, k int) []string {
	t := &trie{}
	for _, w := range words {
		t.insert(w)
	}
	q := []*pair{{str: "", t: t}}
	pq := PriorityQueue{}
	for len(q) > 0 {
		tp := q[0]
		str, t := tp.str, tp.t
		q = q[1:]
		if t.count > 0 {
			heap.Push(&pq, &WordFreq{word: str, freq: t.count})
		}
		for i := 0; i < 26; i++ {
			if t.child[i] == nil {
				continue
			}
			q = append(q, &pair{str: str + string(byte(i)+'a'), t: t.child[i]})
		}
	}
	out := make([]string, 0, k)
	for k > 0 {
		w := heap.Pop(&pq).(*WordFreq).word
		out = append(out, w)
		k--
	}
	return out
}

type pair struct {
	str string
	t   *trie
}

type trie struct {
	child [26]*trie
	count int
}

func (t *trie) insert(word string) {
	for i := range word {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = &trie{}
		}
		t = t.child[idx]
	}
	t.count++
}
