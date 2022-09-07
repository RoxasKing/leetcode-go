package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func uniqueLetterString(s string) int {
	h := [26][]int{}
	for i := 0; i < 26; i++ {
		h[i] = append(h[i], -1)
	}
	for i, c := range s {
		h[c-'A'] = append(h[c-'A'], i)
	}
	for i := 0; i < 26; i++ {
		h[i] = append(h[i], len(s))
	}
	o := 0
	for i := 0; i < 26; i++ {
		for j := 1; j < len(h[i])-1; j++ {
			o += (h[i][j] - h[i][j-1]) * (h[i][j+1] - h[i][j])
		}
	}
	return o
}
