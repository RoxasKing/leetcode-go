package main

// Difficulty:
// Easy

func oddCells(m int, n int, indices [][]int) int {
	rc := make([]int, m)
	cc := make([]int, n)
	for _, e := range indices {
		rc[e[0]]++
		cc[e[1]]++
	}
	x, y := 0, 0
	for _, v := range rc {
		x += v % 2
	}
	for _, v := range cc {
		y += v % 2
	}
	return x*(n-y) + y*(m-x)
}
