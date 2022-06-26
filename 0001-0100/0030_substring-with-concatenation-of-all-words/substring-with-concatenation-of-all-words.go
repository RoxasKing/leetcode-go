package main

// Difficulty:
// Hard

// Tags:
// Sliding Window
// Hash

func findSubstring(s string, words []string) []int {
	m, n, k := len(s), len(words), len(words[0])
	if m < n*k {
		return nil
	}
	var o []int
	for offset := 0; offset < k; offset++ {
		freq := map[string]int{}
		for _, word := range words {
			freq[word]++
		}
		for l, r := offset, offset; r+k <= m; r += k {
			for freq[s[r:r+k]]--; freq[s[r:r+k]] < 0; l += k {
				freq[s[l:l+k]]++
			}
			if r-l+k == n*k {
				o = append(o, l)
			}
		}
	}
	return o
}
