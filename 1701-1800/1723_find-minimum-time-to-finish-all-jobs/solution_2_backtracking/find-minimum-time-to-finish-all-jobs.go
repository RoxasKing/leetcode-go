package main

// Tags:
// Binary Search
// Backtracking

func minimumTimeRequired(jobs []int, k int) int {
	l, r := 1, 12*int(1e7)
	for l < r {
		limit := l + (r-l)>>1
		if !isValid(jobs, k, limit) {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func isValid(jobs []int, k, limit int) bool {
	return backtrack(jobs, make([]int, k), limit, 0)
}

func backtrack(jobs []int, workers []int, limit, idx int) bool {
	if idx == len(jobs) {
		for _, worker := range workers {
			if worker > limit {
				return false
			}
		}
		return true
	}
	for i := range workers {
		if workers[i]+jobs[idx] <= limit {
			workers[i] += jobs[idx]
			if backtrack(jobs, workers, limit, idx+1) {
				return true
			}
			workers[i] -= jobs[idx]
		}
		if workers[i] == 0 || workers[i]+jobs[idx] == limit {
			return false
		}
	}
	return false
}
