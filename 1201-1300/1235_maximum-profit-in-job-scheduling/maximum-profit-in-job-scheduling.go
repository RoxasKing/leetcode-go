package main

import "sort"

// Dynamic Programming

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	set := map[int]struct{}{}
	for i := 0; i < n; i++ {
		set[startTime[i]] = struct{}{}
		set[endTime[i]] = struct{}{}
	}

	times := make([]int, 0, len(set))
	for t := range set {
		times = append(times, t)
	}
	sort.Ints(times)

	dict := map[int]int{}
	for i, point := range times {
		dict[point] = i
	}

	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	m := len(times)
	f := make([]int, m)

	for i := 1; i < m; i++ {
		f[i] = f[i-1]
		for len(jobs) > 0 && jobs[0][1] == times[i] {
			job := jobs[0]
			jobs = jobs[1:]
			f[i] = Max(f[i], f[dict[job[0]]]+job[2])
		}
	}

	return f[m-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
