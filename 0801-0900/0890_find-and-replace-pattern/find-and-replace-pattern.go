package main

// Difficulty:
// Medium

// Tags:
// Hash

func findAndReplacePattern(words []string, pattern string) []string {
	o := []string{}
	for _, w := range words {
		ha, hb := [27]int{}, [27]int{}
		ok := true
		for i := range pattern {
			u, v := int(pattern[i]-'a')+1, int(w[i]-'a')+1
			if ha[u] != 0 && ha[u] != v || hb[v] != 0 && hb[v] != u {
				ok = false
				break
			}
			ha[u], hb[v] = v, u
		}
		if ok {
			o = append(o, w)
		}
	}
	return o
}
