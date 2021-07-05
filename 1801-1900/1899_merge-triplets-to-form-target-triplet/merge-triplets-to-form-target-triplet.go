package main

// Tags:
// Greedy Algorithm

func mergeTriplets(triplets [][]int, target []int) bool {
	t0, t1, t2 := target[0], target[1], target[2]
	a, b, c := 0, 0, 0
	for _, t := range triplets {
		if t[0] <= t0 && t[1] <= t1 && t[2] <= t2 {
			a, b, c = Max(a, t[0]), Max(b, t[1]), Max(c, t[2])
		}
	}
	return a == t0 && b == t1 && c == t2
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
