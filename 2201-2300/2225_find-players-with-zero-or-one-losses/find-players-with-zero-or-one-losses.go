package main

// Difficulty:
// Medium

// Tags:
// Hash

func findWinners(matches [][]int) [][]int {
	d := [1e5 + 1]int{}
	h := [1e5 + 1]bool{}
	for _, e := range matches {
		h[e[0]] = true
		h[e[1]] = true
		d[e[1]]++
	}
	o := make([][]int, 2)
	for i := 1; i <= 1e5; i++ {
		if !h[i] {
			continue
		}
		if d[i] == 0 {
			o[0] = append(o[0], i)
		} else if d[i] == 1 {
			o[1] = append(o[1], i)
		}
	}
	return o
}
