package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func findLUSlength(strs []string) int {
	lcs := func(a, b string) bool {
		i := 0
		for j := 0; i < len(a) && j < len(b); j++ {
			if a[i] == b[j] {
				i++
			}
		}
		return i == len(a)
	}
	o := -1
	for i, a := range strs {
		c := 0
		for j, b := range strs {
			if i != j && len(a) <= len(b) && lcs(a, b) {
				break
			}
			c++
		}
		if c == len(strs) && o < len(a) {
			o = len(a)
		}
	}
	return o
}
