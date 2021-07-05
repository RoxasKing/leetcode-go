package main

func numEquivDominoPairs(dominoes [][]int) int {
	count := make([]int, 100)
	for _, d := range dominoes {
		if d[0] > d[1] {
			d[0], d[1] = d[1], d[0]
		}
		count[d[0]*10+d[1]]++
	}
	out := 0
	for _, c := range count {
		out += c * (c - 1) / 2
	}
	return out
}
