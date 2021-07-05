package main

import "sort"

// Trie
// Offline Algorithm
// Bit Manipulation

func maximizeXor(nums []int, queries [][]int) []int {
	sort.Ints(nums)
	m, n := len(nums), len(queries)
	qs := make([][3]int, n)
	for i := range qs {
		qs[i] = [3]int{i, queries[i][0], queries[i][1]}
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i][2] < qs[j][2] })
	t := &Trie{}
	out := make([]int, n)
	for i, j := 0, 0; i < n; i++ {
		q := qs[i]
		idx, xi, mi := q[0], q[1], q[2]
		for ; j < m && nums[j] <= mi; j++ {
			t.Insert(nums[j])
		}
		if j == 0 {
			out[idx] = -1
		} else {
			out[idx] = t.Query(xi)
		}
	}
	return out
}

type Trie struct {
	child [2]*Trie
}

func (t *Trie) Insert(num int) {
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
}

func (t *Trie) Query(num int) int {
	out := 0
	for i := 31; i >= 0; i-- {
		idx := (num >> i) & 1
		if t.child[1-idx] != nil {
			out |= 1 << i
			t = t.child[1-idx]
		} else {
			t = t.child[idx]
		}
	}
	return out
}
