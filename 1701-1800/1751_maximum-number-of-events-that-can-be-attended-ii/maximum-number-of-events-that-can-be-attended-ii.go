package main

import "sort"

// Discretization
// Dynamic Programming

func maxValue(events [][]int, k int) int {
	if k == 1 {
		out := 0
		for _, e := range events {
			out = Max(out, e[2])
		}
		return out
	}

	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })

	set := map[int]struct{}{}
	for _, event := range events {
		set[event[0]] = struct{}{}
		set[event[1]] = struct{}{}
	}

	pts := make([]int, 0, len(set))
	for pt := range set {
		pts = append(pts, pt)
	}
	sort.Ints(pts)

	dict := make(map[int]int)
	for i, pt := range pts {
		dict[pt] = i + 1
	}

	n := len(pts)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		copy(f[i], f[i-1])
		for ; len(events) > 0 && dict[events[0][1]] <= i; events = events[1:] {
			st, val := events[0][0], events[0][2]
			for j := 0; j < k; j++ {
				f[i][j+1] = Max(f[i][j+1], f[dict[st]-1][j]+val)
			}
		}
	}

	return f[n][k]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
