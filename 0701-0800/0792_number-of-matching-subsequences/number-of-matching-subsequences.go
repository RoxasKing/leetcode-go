package main

// Difficulty:
// Medium

// Tags:
// Hash

func numMatchingSubseq(s string, words []string) int {
	type pair struct{ id, i int }
	h := [26][]pair{}
	for id, w := range words {
		b := int(w[0] - 'a')
		h[b] = append(h[b], pair{id, 0})
	}
	o := 0
	for _, c := range s {
		b := int(c - 'a')
		n := len(h[b])
		for _, p := range h[b] {
			if p.i++; p.i == len(words[p.id]) {
				o++
				continue
			}
			nb := int(words[p.id][p.i] - 'a')
			h[nb] = append(h[nb], p)
		}
		h[b] = h[b][n:]
	}
	return o
}
