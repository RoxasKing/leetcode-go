package main

import "sort"

// Greedy Algorithm
func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, task := range tasks {
		count[task-'A']++
	}
	sort.Ints(count)
	gap := count[25] - 1
	slot := gap * n
	for i := 24; i >= 0 && count[i] > 0 && slot > 0; i-- {
		slot -= Min(count[i], gap)
	}
	if slot > 0 {
		return slot + len(tasks)
	}
	return len(tasks)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
