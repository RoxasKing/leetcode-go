package main

// Difficulty:
// Hard

// Tags:
// Topological Sort
// Hash

func alienOrder(words []string) string {
	h := [26]bool{}
	g := [26][26]bool{}
	d := [26]int{}
	cnt := 0
	for k, a := range words {
		for _, c := range a {
			if !h[c-'a'] {
				h[c-'a'] = true
				cnt++
			}
		}
		for _, b := range words[k:] {
			ok := false
			for i := 0; i < min(len(a), len(b)); i++ {
				if a[i] != b[i] {
					u, v := int(a[i]-'a'), int(b[i]-'a')
					if !g[u][v] {
						g[u][v] = true
						d[v]++
					}
					ok = true
					break
				}
			}
			if !ok && len(a) > len(b) {
				return ""
			}
		}
	}
	q := []int{}
	for i := 0; i < 26; i++ {
		if h[i] && d[i] == 0 {
			q = append(q, i)
		}
	}
	o := make([]byte, 0, cnt)
	for ; len(q) > 0; q = q[1:] {
		if len(o) == cnt {
			return ""
		}
		u := q[0]
		o = append(o, byte(u)+'a')
		for v := 0; v < 26; v++ {
			if !g[u][v] {
				continue
			}
			if d[v]--; d[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(o) < cnt {
		return ""
	}
	return string(o)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
