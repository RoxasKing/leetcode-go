package main

func countBalls(lowLimit int, highLimit int) int {
	cnt := [46]int{}
	for i := lowLimit; i <= highLimit; i++ {
		idx := 0
		for j := i; j > 0; j /= 10 {
			idx += j % 10
		}
		cnt[idx]++
	}
	out := 0
	for i := 1; i <= 45; i++ {
		out = Max(out, cnt[i])
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
