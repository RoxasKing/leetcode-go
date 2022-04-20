package main

// Difficulty:
// Easy

// Tags:
// Hash

func mostCommonWord(paragraph string, banned []string) string {
	h := map[string]struct{}{}
	for _, s := range banned {
		h[s] = struct{}{}
	}
	c := map[string]int{"": 0}
	o := ""
	n := len(paragraph)
	for i := 0; i < n; {
		for ; i < n && !isLetter(paragraph[i]); i++ {
		}
		if i == n {
			break
		}
		arr := []byte{}
		for ; i < n && isLetter(paragraph[i]); i++ {
			c := paragraph[i]
			if 'A' <= c && c <= 'Z' {
				c += 'a' - 'A'
			}
			arr = append(arr, c)
		}
		s := string(arr)
		if _, ok := h[s]; ok {
			continue
		}
		if c[s]++; c[o] < c[s] {
			o = s
		}
	}
	return o
}

func isLetter(c byte) bool { return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' }
