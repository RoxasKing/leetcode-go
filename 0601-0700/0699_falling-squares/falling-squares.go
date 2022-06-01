package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Segment Tree
// Discretization
// Hash

func fallingSquares(positions [][]int) []int {
	set := map[int]struct{}{}
	a := []int{}
	for _, p := range positions {
		x, y := p[0], p[0]+p[1]-1
		if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			a = append(a, x)
		}
		if _, ok := set[y]; !ok {
			set[y] = struct{}{}
			a = append(a, y)
		}
	}
	sort.Ints(a)
	h := map[int]int{}
	for i, x := range a {
		h[x] = i
	}
	n := len(a)
	f := make([]int, n*4)
	b := make([]int, n*4)
	s, t := 0, n-1
	var query func(i, s, t, l, r int) int
	query = func(i, s, t, l, r int) int {
		if r < s || t < l {
			return 0
		}
		if l <= s && t <= r {
			return f[i]
		}
		if b[i] != 0 && s != t {
			f[i*2], f[i*2+1] = b[i], b[i]
			b[i*2], b[i*2+1] = b[i], b[i]
			b[i] = 0
		}
		m := (s + t) / 2
		return max(query(i*2, s, m, l, r), query(i*2+1, m+1, t, l, r))
	}
	var update func(i, s, t, l, r, v int)
	update = func(i, s, t, l, r, v int) {
		if r < s || t < l {
			return
		}
		if l <= s && t <= r {
			f[i], b[i] = v, v
			return
		}
		if b[i] != 0 && s != t {
			f[i*2], f[i*2+1] = b[i], b[i]
			b[i*2], b[i*2+1] = b[i], b[i]
			b[i] = 0
		}
		m := (s + t) / 2
		update(i*2, s, m, l, r, v)
		update(i*2+1, m+1, t, l, r, v)
		f[i] = max(f[i], max(f[i*2], f[i*2+1]))
	}
	max := 0
	o := make([]int, len(positions))
	for i, p := range positions {
		l, r := h[p[0]], h[p[0]+p[1]-1]
		v := query(1, s, t, l, r)
		if v += p[1]; max < v {
			max = v
		}
		o[i] = max
		update(1, s, t, l, r, v)
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
