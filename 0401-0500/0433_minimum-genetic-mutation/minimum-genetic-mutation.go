package main

// Difficulty:
// Medium

// Tags:
// Hash
// BFS

func minMutation(start string, end string, bank []string) int {
	h := map[string]struct{}{}
	for _, s := range bank {
		h[s] = struct{}{}
	}
	if _, ok := h[end]; !ok {
		return -1
	}
	type pair struct {
		k string
		v int
	}
	out := -1
	for q := []pair{{start, 0}}; len(q) > 0; q = q[1:] {
		p := q[0]
		if p.k == end {
			out = p.v
			break
		}
		t := []byte(p.k)
		for i := 0; i < len(t); i++ {
			for c := byte('A'); c <= byte('Z'); c++ {
				if c == p.k[i] {
					continue
				}
				t[i] = c
				if _, ok := h[string(t)]; !ok {
					t[i] = p.k[i]
					continue
				}
				q = append(q, pair{string(t), p.v + 1})
				t[i] = p.k[i]
			}
		}
	}
	return out
}
