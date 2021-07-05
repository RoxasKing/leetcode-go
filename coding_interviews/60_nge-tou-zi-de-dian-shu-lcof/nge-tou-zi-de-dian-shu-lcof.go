package main

// Tags:
// Dynamic Programming
func dicesProbability(n int) []float64 {
	cnt, tmp := make([]int, 6*n+1), make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		cnt[i] = 1
	}
	t := 6
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			tmp[j] = 0
		}
		for j := i; j <= 6*i; j++ {
			tmp[j] = 0
			for k := 1; k <= 6 && k < j; k++ {
				tmp[j] += cnt[j-k]
			}
		}
		cnt, tmp = tmp, cnt
		t *= 6
	}
	cnt = cnt[n:]
	out := make([]float64, len(cnt))
	for i := range out {
		out[i] = float64(cnt[i]) / float64(t)
	}
	return out
}
