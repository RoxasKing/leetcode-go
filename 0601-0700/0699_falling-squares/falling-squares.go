package main

import (
	"sort"
)

// Tags:
// Segment Tree
func fallingSquares(positions [][]int) []int {
	set := make(map[int]struct{})
	for _, pos := range positions {
		set[pos[0]] = struct{}{}
		set[pos[0]+pos[1]-1] = struct{}{}
	}
	idxs := []int{}
	for pos := range set {
		idxs = append(idxs, pos)
	}
	sort.Ints(idxs)
	dict := make(map[int]int)
	for i, pos := range idxs {
		dict[pos] = i
	}
	st := NewSegmentTree(len(idxs))

	n := len(positions)
	out := make([]int, n)
	for i, pos := range positions {
		l, r := dict[pos[0]], dict[pos[0]+pos[1]-1]
		base := st.Query(1, st.s, st.t, l, r)
		st.Update(1, st.s, st.t, l, r, base+pos[1])
		out[i] = st.Query(1, st.s, st.t, st.s, st.t)
	}
	return out
}

type SegmentTree struct {
	f, b []int
	s, t int
}

func NewSegmentTree(n int) SegmentTree {
	f := make([]int, 4*n)
	b := make([]int, 4*n)
	s, t := 0, n-1
	return SegmentTree{
		f: f,
		b: b,
		s: s,
		t: t,
	}
}

func (st *SegmentTree) Query(idx, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		return st.f[idx]
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 {
		st.f[idx<<1] = st.b[idx]
		st.f[idx<<1+1] = st.b[idx]
		st.b[idx<<1] = st.b[idx]
		st.b[idx<<1+1] = st.b[idx]
		st.b[idx] = 0
	}
	return Max(st.Query(idx<<1, s, m, l, r), st.Query(idx<<1+1, m+1, t, l, r))
}

func (st *SegmentTree) Update(idx, s, t, l, r, val int) {
	if t < l || s > r {
		return
	}
	if l <= s && t <= r {
		st.f[idx] = val
		st.b[idx] = val
		return
	}
	m := (s + t) >> 1
	if st.b[idx] != 0 && s != t {
		st.f[idx<<1] = st.b[idx]
		st.f[idx<<1+1] = st.b[idx]
		st.b[idx<<1] = st.b[idx]
		st.b[idx<<1+1] = st.b[idx]
		st.b[idx] = 0
	}
	st.Update(idx<<1, s, m, l, r, val)
	st.Update(idx<<1+1, m+1, t, l, r, val)
	st.f[idx] = Max(st.f[idx], Max(st.f[idx<<1], st.f[idx<<1+1]))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
