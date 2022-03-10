package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Prefix Sum
// Binary Search

func platesBetweenCandles(s string, queries [][]int) []int {
	pSum := [][2]int{{-1, 0}}
	for cnt, ppl, i := 0, -1, 0; i < len(s); i++ {
		if s[i] == '*' {
			cnt++
			continue
		}
		if ppl < 0 {
			cnt = 0
		}
		pSum = append(pSum, [2]int{i, cnt})
		cnt, ppl = 0, i
	}
	for i := 1; i < len(pSum); i++ {
		pSum[i][1] += pSum[i-1][1]
	}
	out := make([]int, len(queries))
	for k := range queries {
		l, r := queries[k][0], queries[k][1]
		i := sort.Search(len(pSum), func(i int) bool { return pSum[i][0] >= l })
		j := sort.Search(len(pSum)-i, func(j int) bool { return pSum[i+j][0] >= r }) + i
		if s[r] == '*' {
			j--
		}
		if i > j {
			continue
		}
		out[k] = pSum[j][1] - pSum[i][1]
	}
	return out
}
