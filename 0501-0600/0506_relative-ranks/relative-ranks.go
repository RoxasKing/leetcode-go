package main

import (
	"sort"
	"strconv"
)

// Difficulty:
// Easy

// Tags:
// Sort

func findRelativeRanks(score []int) []string {
	n := len(score)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool { return score[idx[i]] > score[idx[j]] })
	out := make([]string, n)
	out[idx[0]] = "Gold Medal"
	if n > 1 {
		out[idx[1]] = "Silver Medal"
	}
	if n > 2 {
		out[idx[2]] = "Bronze Medal"
	}
	for i := 3; i < n; i++ {
		out[idx[i]] = strconv.Itoa(i + 1)
	}
	return out
}
