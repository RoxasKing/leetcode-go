package main

// Difficulty:
// Medium

func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := (m + n) * mean
	for _, x := range rolls {
		sum -= x
	}
	avg := sum / n
	rem := sum - avg*n
	if avg < 1 || avg > 6 || avg == 6 && rem > 0 {
		return []int{}
	}
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = avg
		if rem > 0 {
			rem--
			out[i]++
		}
	}
	return out
}
