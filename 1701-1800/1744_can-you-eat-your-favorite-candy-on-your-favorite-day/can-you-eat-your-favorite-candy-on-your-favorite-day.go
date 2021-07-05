package main

// Tags:
// Prefix Sum
// Math

func canEat(candiesCount []int, queries [][]int) []bool {
	n, m := len(candiesCount), len(queries)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + candiesCount[i]
	}

	out := make([]bool, m)
	for i, q := range queries {
		t, d, c := q[0], q[1], q[2]
		out[i] = pSum[t]/c-1 < d && d < pSum[t+1]
	}
	return out
}
