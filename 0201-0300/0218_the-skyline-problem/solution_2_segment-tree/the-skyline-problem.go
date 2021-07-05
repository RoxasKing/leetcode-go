package main

import (
	"sort"
)

// Tags:
// Segment Tree
func getSkyline(buildings [][]int) [][]int {
	set := make(map[int]struct{})
	for _, building := range buildings {
		l, r := building[0], building[1]
		set[l] = struct{}{}
		set[r] = struct{}{}
	}
	arr := make([]int, 0, len(set))
	for building := range set {
		arr = append(arr, building)
	}
	sort.Ints(arr)
	dict := make(map[int]int)
	for i, building := range arr {
		dict[building] = i
	}
	n := len(arr)

	st := NewSegmentTree(n)
	for _, building := range buildings {
		l, r, val := dict[building[0]], dict[building[1]]-1, building[2]
		st.Update(1, st.s, st.t, l, r, val)
	}

	out := [][]int{}
	for _, b := range arr {
		idx := dict[b]
		h := st.Query(1, st.s, st.t, idx, idx)
		if len(out) == 0 || out[len(out)-1][1] != h {
			out = append(out, []int{b, h})
		}
	}
	return out
}

type SegmentTree struct {
	f, b []int
	s, t int
}

func NewSegmentTree(n int) SegmentTree {
	return SegmentTree{
		f: make([]int, 4*n),
		b: make([]int, 4*n),
		s: 0,
		t: n - 1,
	}
}

func (st *SegmentTree) Query(idx, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		if st.b[idx] != 0 {
			st.f[idx] = Max(st.f[idx], st.b[idx])
			st.b[idx] = 0
		}
		return st.f[idx]
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 && s != t {
		st.f[idx<<1] = Max(st.f[idx<<1], st.b[idx])
		st.f[idx<<1+1] = Max(st.f[idx<<1+1], st.b[idx])
		st.b[idx<<1] = Max(st.b[idx<<1], st.b[idx])
		st.b[idx<<1+1] = Max(st.b[idx<<1+1], st.b[idx])
		st.b[idx] = 0
	}
	return Max(st.Query(idx<<1, s, m, l, r), st.Query(idx<<1+1, m+1, t, l, r))
}

func (st *SegmentTree) Update(idx, s, t, l, r, val int) {
	if t < l || s > r {
		return
	}
	if l <= s && t <= r {
		if st.b[idx] != 0 {
			st.f[idx] = Max(st.f[idx], st.b[idx])
		}
		st.f[idx] = Max(st.f[idx], val)
		st.b[idx] = st.f[idx]
		return
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 && s != t {
		st.f[idx<<1] = Max(st.f[idx<<1], st.b[idx])
		st.f[idx<<1+1] = Max(st.f[idx<<1+1], st.b[idx])
		st.b[idx<<1] = st.f[idx<<1]
		st.b[idx<<1+1] = st.f[idx<<1+1]
		st.b[idx] = 0
	}
	st.Update(idx<<1, s, m, l, r, val)
	st.Update(idx<<1+1, m+1, t, l, r, val)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
