package main

// Tags:
// Two Pointers + Greedy Algorithm
func partitionLabels(S string) []int {
	last := [26]int{}
	for i := range S {
		last[S[i]-'a'] = i
	}

	out := []int{}
	for l, r, i := 0, 0, 0; i < len(S); i++ {
		r = Max(r, last[S[i]-'a'])
		if i == r {
			out = append(out, r+1-l)
			l = r + 1
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
