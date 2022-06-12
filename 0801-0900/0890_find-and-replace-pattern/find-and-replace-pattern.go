package main

// Difficulty:
// Medium

// Tags:
// Hash

func findAndReplacePattern(words []string, pattern string) []string {
	o := []string{}
	for _, word := range words {
		hu, hv := [26]int{}, [26]int{}
		for i := 0; i < 26; i++ {
			hu[i], hv[i] = -1, -1
		}
		ok := true
		for i := range pattern {
			u, v := int(pattern[i]-'a'), int(word[i]-'a')
			if hu[u] == -1 && hv[v] == -1 {
				hu[u], hv[v] = v, u
			} else if hu[u] != v || hv[v] != u {
				ok = false
				break
			}
		}
		if ok {
			o = append(o, word)
		}
	}
	return o
}
