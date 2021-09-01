package main

import "sort"

// Tags:
// Discretization
// Dynamic Programming

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	m := len(startTime)
	jobs := make([]*Job, m)
	set := map[int]struct{}{}
	for i := 0; i < m; i++ {
		st, ed := startTime[i], endTime[i]
		jobs[i] = &Job{st: st, ed: ed, pf: profit[i]}
		set[st] = struct{}{}
		set[ed] = struct{}{}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].ed < jobs[j].ed })

	n := len(set)
	ts := make([]int, 0, n)
	for t := range set {
		ts = append(ts, t)
	}
	sort.Ints(ts)
	dict := map[int]int{}
	for i, t := range ts {
		dict[t] = i + 1
	}

	f := make([]int, n+1)
	for i, j := 1, 0; i <= n; i++ {
		f[i] = Max(f[i], f[i-1])
		for ; j < m && jobs[j].ed <= ts[i-1]; j++ {
			f[i] = Max(f[i], f[dict[jobs[j].st]]+jobs[j].pf)
		}
	}
	return f[n]
}

type Job struct {
	st, ed, pf int
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
