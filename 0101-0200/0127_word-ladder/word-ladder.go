package main

// Difficulty:
// Hard

// Tags:
// Hash
// BFS

func ladderLength(beginWord string, endWord string, wordList []string) int {
	h := map[string]struct{}{}
	for _, w := range wordList {
		h[w] = struct{}{}
	}
	if _, ok := h[endWord]; !ok {
		return 0
	}
	delete(h, endWord)
	src, dst := map[string]struct{}{beginWord: {}}, map[string]struct{}{endWord: {}}
	out := 1
	for len(src) != 0 && len(dst) != 0 {
		out++
		if len(src) > len(dst) {
			src, dst = dst, src
		}
		for w := range src {
			delete(h, w)
		}
		tmp := map[string]struct{}{}
		for w := range src {
			letters := []byte(w)
			for i := range letters {
				for k := 0; k < 26; k++ {
					letters[i] = byte(k) + 'a'
					nxt := string(letters)
					if _, ok := dst[nxt]; ok {
						return out
					}
					if _, ok := h[nxt]; ok {
						tmp[nxt] = struct{}{}
					}
				}
				letters[i] = w[i]
			}
		}
		src = tmp
	}
	return 0
}
