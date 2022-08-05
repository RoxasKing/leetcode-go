package main

// Difficulty:
// Medium

// Tags:
// Hash

func wordSubsets(words1 []string, words2 []string) []string {
	h := [26]int{}
	for _, w := range words2 {
		ht := [26]int{}
		for _, c := range w {
			i := int(c - 'a')
			if ht[i]++; h[i] < ht[i] {
				h[i] = ht[i]
			}
		}
	}
	o := []string{}
	for _, w := range words1 {
		ht := [26]int{}
		for _, c := range w {
			ht[c-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if h[i] > ht[i] {
				ok = false
				break
			}
		}
		if ok {
			o = append(o, w)
		}
	}
	return o
}
