package main

// Difficulty:
// Hard

// Tags:
// Greedy

func findMinMoves(machines []int) int {
	n := len(machines)
	tot := 0
	for _, machine := range machines {
		tot += machine
	}
	if tot%n != 0 {
		return -1
	}
	avg := tot / n
	sum := 0
	out := 0
	for _, num := range machines {
		num -= avg
		sum += num
		out = Max(out, Max(num, Abs(sum)))
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
