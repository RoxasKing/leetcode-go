package main

// Difficulty:
// Medium

// Tags:
// Hash
// Two Pointers
// Greedy

func partitionLabels(s string) []int {
	last := [26]int{}
	for i := range s {
		last[s[i]-'a'] = i
	}
	out := []int{}
	for l, r, i := 0, 0, 0; i < len(s); i++ {
		r = Max(r, last[s[i]-'a'])
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
