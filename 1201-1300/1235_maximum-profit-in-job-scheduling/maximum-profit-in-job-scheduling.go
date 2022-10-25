package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Discretization
// Dynamic Programming

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	m := len(startTime)
	jobs := make([]Job, m)
	set := map[int]struct{}{}
	for i := 0; i < m; i++ {
		st, ed := startTime[i], endTime[i]
		jobs[i] = Job{st, ed, profit[i]}
		set[st] = struct{}{}
		set[ed] = struct{}{}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].ed < jobs[j].ed })
	n := len(set)
	a := make([]int, 0, n)
	for t := range set {
		a = append(a, t)
	}
	sort.Ints(a)
	h := map[int]int{}
	for i, t := range a {
		h[t] = i + 1
	}
	f := make([]int, n+1)
	for i, j := 0, 0; i < n; i++ {
		f[i+1] = max(f[i+1], f[i])
		for ; j < m && jobs[j].ed <= a[i]; j++ {
			f[i+1] = max(f[i+1], f[h[jobs[j].st]]+jobs[j].pf)
		}
	}
	return f[n]
}

type Job struct{ st, ed, pf int }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
