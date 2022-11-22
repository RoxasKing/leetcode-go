package main

// Difficulty:
// Easy

// Tags:
// Hash

func countConsistentStrings(allowed string, words []string) int {
	h := [26]bool{}
	for _, c := range allowed {
		h[c-'a'] = true
	}
	o := 0
	for _, w := range words {
		ok := true
		for _, c := range w {
			if !h[c-'a'] {
				ok = false
				break
			}
		}
		if ok {
			o++
		}
	}
	return o
}
