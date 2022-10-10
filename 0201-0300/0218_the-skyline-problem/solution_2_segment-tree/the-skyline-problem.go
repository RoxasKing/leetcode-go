package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Discretization
// Hash
// Segment Tree

func getSkyline(buildings [][]int) [][]int {
	set := map[int]struct{}{}
	for _, building := range buildings {
		l, r := building[0], building[1]
		set[l] = struct{}{}
		set[r] = struct{}{}
	}
	arr := make([]int, 0, len(set))
	for x := range set {
		arr = append(arr, x)
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
	var update func(i, s, t, l, r, v int)
	update = func(i, s, t, l, r, v int) {
		if r < s || t < l {
			return
		}
		if l <= s && t <= r {
			f[i] = max(f[i], v)
			b[i] = f[i]
			return
		}
		if b[i] != 0 && s != t {
			f[i*2], f[i*2+1] = max(f[i*2], b[i]), max(f[i*2+1], b[i])
			b[i*2], b[i*2+1] = f[i*2], f[i*2+1]
			b[i] = 0
		}
		m := (s + t) / 2
		update(i*2, s, m, l, r, v)
		update(i*2+1, m+1, t, l, r, v)
	}
	var query func(i, s, t, l, r int) int
	query = func(i, s, t, l, r int) int {
		if r < s || t < l {
			return 0
		}
		if l <= s && t <= r {
			if b[i] != 0 {
				f[i] = b[i]
			}
			return f[i]
		}
		if b[i] != 0 && s != t {
			f[i*2], f[i*2+1] = max(f[i*2], b[i]), max(f[i*2+1], b[i])
			b[i*2], b[i*2+1] = f[i*2], f[i*2+1]
			b[i] = 0
		}
		m := (s + t) / 2
		return max(query(i*2, s, m, l, r), query(i*2+1, m+1, t, l, r))
	}
	for _, building := range buildings {
		l, r, v := h[building[0]], h[building[1]]-1, building[2]
		update(1, s, t, l, r, v)
	}
	o := [][]int{}
	for _, x := range arr {
		i := h[x]
		v := query(1, s, t, i, i)
		if len(o) == 0 || o[len(o)-1][1] != v {
			o = append(o, []int{x, v})
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
