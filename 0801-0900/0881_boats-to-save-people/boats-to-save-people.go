package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Greedy
// Two Pointers

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	out := 0
	for l, r := 0, len(people)-1; l <= r; r-- {
		if people[l]+people[r] <= limit {
			l++
		}
		out++
	}
	return out
}
