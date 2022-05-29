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
	arr := []int{}
	for _, pos := range positions {
		a, b := pos[0], pos[0]+pos[1]-1
		if _, ok := set[a]; !ok {
			set[a] = struct{}{}
			arr = append(arr, a)
		}
		if _, ok := set[b]; !ok {
			set[b] = struct{}{}
			arr = append(arr, b)
		}
	}
	sort.Ints(arr)
	h := map[int]int{}
	for i, v := range arr {
		h[v] = i
	}
	n := len(arr)
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
		m := (s + t) / 2
		if b[i] != 0 {
			f[i*2] = b[i]
			f[i*2+1] = b[i]
			b[i*2] = b[i]
			b[i*2+1] = b[i]
			b[i] = 0
		}
		return max(query(i*2, s, m, l, r), query(i*2+1, m+1, t, l, r))
	}
	var update func(i, s, t, l, r, v int)
	update = func(i, s, t, l, r, v int) {
		if r < s || t < l {
			return
		}
		if l <= s && t <= r {
			f[i] = v
			b[i] = v
			return
		}
		m := (s + t) / 2
		if b[i] != 0 && s != t {
			f[i*2] = b[i]
			f[i*2+1] = b[i]
			b[i*2] = b[i]
			b[i*2+1] = b[i]
			b[i] = 0
		}
		update(i*2, s, m, l, r, v)
		update(i*2+1, m+1, t, l, r, v)
		f[i] = max(f[i], max(f[i*2], f[i*2+1]))
	}
	o := make([]int, len(positions))
	tall := 0
	for i, p := range positions {
		l, r := h[p[0]], h[p[0]+p[1]-1]
		v := query(1, s, t, l, r)
		v += p[1]
		update(1, s, t, l, r, v)
		if tall < v {
			tall = v
		}
		o[i] = tall
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
