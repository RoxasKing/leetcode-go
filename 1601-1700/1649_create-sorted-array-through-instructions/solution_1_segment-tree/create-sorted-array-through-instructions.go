package main

import "sort"

// Segment Tree

func createSortedArray(instructions []int) int {
	set := map[int]struct{}{}
	arr := []int{}
	for _, num := range instructions {
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
			arr = append(arr, num)
		}
	}
	sort.Ints(arr)

	dict := map[int]int{}
	for i, num := range arr {
		dict[num] = i
	}

	n := len(arr)
	f := make([]int, 4*n)

	out := 0
	for _, num := range instructions {
		idx := dict[num]
		if 0 < idx && idx < n-1 {
			out += Min(query(f, 1, 0, n-1, 0, idx-1), query(f, 1, 0, n-1, idx+1, n-1))
		}
		add(f, 1, 0, n-1, idx)
	}
	return out % (1e9 + 7)
}

func add(f []int, i, s, t, idx int) {
	if t < idx || s > idx {
		return
	}
	if s == idx && t == idx {
		f[i]++
		return
	}
	m := (s + t) >> 1
	add(f, i<<1, s, m, idx)
	add(f, i<<1+1, m+1, t, idx)
	f[i] = f[i<<1] + f[i<<1+1]
}

func query(f []int, i, s, t, l, r int) int {
	if t < l || s > r {
		return 0
	}
	if l <= s && t <= r {
		return f[i]
	}
	m := (s + t) >> 1
	return query(f, i<<1, s, m, l, r) + query(f, i<<1+1, m+1, t, l, r)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
