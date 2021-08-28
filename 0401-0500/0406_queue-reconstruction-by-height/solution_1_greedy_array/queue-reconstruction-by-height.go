package main

import "sort"

// Tags:
// Greedy
// Array

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] > people[j][0]
		}
		return people[i][1] < people[j][1]
	})
	out := make([][]int, len(people))
	for _, p := range people {
		copy(out[p[1]+1:], out[p[1]:])
		out[p[1]] = p
	}
	return out
}
