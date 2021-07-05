package main

import "sort"

// Backtracking
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sort.Ints(jobs)
	max, sum := jobs[n-1], 0
	for _, job := range jobs {
		sum += job
	}

	l, r := max, sum
	for l < r {
		maxWorkTime := (l + r) >> 1
		if !valid(jobs, n, k, maxWorkTime) {
			l = maxWorkTime + 1
		} else {
			r = maxWorkTime
		}
	}
	return l
}

func valid(jobs []int, n, k, maxWorkTime int) bool {
	return backtrack(jobs, n, maxWorkTime, 0, make([]int, k))
}

func backtrack(jobs []int, n, maxWorkTime, idx int, worker []int) bool {
	if idx == n {
		return true
	}

	for i := range worker {
		if worker[i]+jobs[idx] <= maxWorkTime {
			worker[i] += jobs[idx]
			if backtrack(jobs, n, maxWorkTime, idx+1, worker) {
				return true
			}
			worker[i] -= jobs[idx]
		}
		if worker[i] == 0 || worker[i]+jobs[idx] == maxWorkTime {
			return false
		}
	}

	return false
}
