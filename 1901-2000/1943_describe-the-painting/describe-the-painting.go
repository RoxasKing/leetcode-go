package main

import "sort"

// Tags:
// Segment Tree
// Discretization

func splitPainting(segments [][]int) [][]int64 {
	set := map[int]struct{}{}
	for _, segment := range segments {
		set[segment[0]] = struct{}{}
		set[segment[1]] = struct{}{}
	}
	points := make([]int, 0, len(set))
	for point := range set {
		points = append(points, point)
	}
	sort.Ints(points)
	dict := map[int]int{}
	for i, p := range points {
		dict[p] = i
	}

	n := len(points)
	f := make([]int, 4*n)
	b := make([]int, 4*n)
	s, t := 0, n-1
	for _, segment := range segments {
		l, r, v := dict[segment[0]], dict[segment[1]]-1, segment[2]
		update(f, b, 1, s, t, l, r, v)
	}

	out := make([][]int64, 0, n-1)
	for i := 0; i < n-1; i++ {
		l, r := dict[points[i]], dict[points[i+1]]-1
		v := query(f, b, 1, s, t, l, r)
		if v == 0 {
			continue
		}
		out = append(out, []int64{int64(points[i]), int64(points[i+1]), int64(v)})
	}
	return out
}

func update(f, b []int, i, s, t, l, r, v int) {
	if t < l || s > r {
		return
	}
	if l <= s && t <= r {
		f[i] += v
		b[i] += v
		return
	}
	m := (s + t) >> 1
	if b[i] != 0 && s != t {
		f[i<<1] += b[i]
		f[i<<1+1] += b[i]
		b[i<<1] += b[i]
		b[i<<1+1] += b[i]
		b[i] = 0
	}
	update(f, b, i<<1, s, m, l, r, v)
	update(f, b, i<<1+1, m+1, t, l, r, v)
}

func query(f, b []int, i, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		return f[i]
	}
	m := (s + t) >> 1
	if b[i] != 0 && s != t {
		f[i<<1] += b[i]
		f[i<<1+1] += b[i]
		b[i<<1] += b[i]
		b[i<<1+1] += b[i]
		b[i] = 0
	}
	return query(f, b, i<<1, s, m, l, r) + query(f, b, i<<1+1, m+1, t, l, r)
}
