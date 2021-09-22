package main

import "sort"

// Tags:
// Greedy
// Sort
// Prefix Sum

func maxmiumScore(cards []int, cnt int) int {
	odd := []int{}
	eve := []int{}
	for _, c := range cards {
		if c&1 == 1 {
			odd = append(odd, c)
		} else {
			eve = append(eve, c)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))
	sort.Sort(sort.Reverse(sort.IntSlice(eve)))
	m, n := len(odd), len(eve)
	pO := make([]int, m+1)
	pE := make([]int, n+1)
	for i := 0; i < m; i++ {
		pO[i+1] = pO[i] + odd[i]
	}
	for i := 0; i < n; i++ {
		pE[i+1] = pE[i] + eve[i]
	}
	out := 0
	for i := 0; i <= m && i <= cnt; i += 2 {
		if cnt-i <= n && pO[i]+pE[cnt-i] > out {
			out = pO[i] + pE[cnt-i]
		}
	}
	return out
}
