package main

// Difficulty:
// Hard

// Tags:
// Hash
// BFS
// Backtracking

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	listSet := map[string]struct{}{}
	for _, w := range wordList {
		listSet[w] = struct{}{}
	}
	if _, ok := listSet[endWord]; !ok {
		return nil
	}
	delete(listSet, endWord)
	h := map[string][]string{}
	src, dst := map[string]struct{}{beginWord: {}}, map[string]struct{}{endWord: {}}
	find, reverse := false, false
	parents := map[string][]string{}
	midSet := map[string]struct{}{}
	for len(src) > 0 && len(dst) > 0 && !find {
		if len(src) > len(dst) {
			src, dst = dst, src
			reverse = !reverse
		}
		for w := range src {
			delete(listSet, w)
		}
		tmp := map[string]struct{}{}
		for w := range src {
			a := []rune(w)
			for i := range a {
				for c := 'a'; c <= 'z'; c++ {
					a[i] = c
					s := string(a)
					if _, ok := dst[s]; ok {
						if reverse {
							h[s] = append(h[s], w)
						} else {
							h[w] = append(h[w], s)
						}
						midSet[s] = struct{}{}
						midSet[w] = struct{}{}
						find = true
						continue
					}
					if _, ok := listSet[s]; ok {
						tmp[s] = struct{}{}
						if reverse {
							h[s] = append(h[s], w)
						} else {
							h[w] = append(h[w], s)
						}
						parents[s] = append(parents[s], w)
					}
				}
				a[i] = rune(w[i])
			}
		}
		src = tmp
	}
	if !find {
		return nil
	}
	allSet := map[string]struct{}{}
	for s := range midSet {
		for q := []string{s}; len(q) > 0; q = q[1:] {
			w := q[0]
			allSet[w] = struct{}{}
			if ps, ok := parents[w]; ok {
				q = append(q, ps...)
			}
		}
	}
	o := [][]string{}
	a := []string{beginWord}
	var backtrack func(cur string)
	backtrack = func(cur string) {
		if cur == endWord {
			t := make([]string, len(a))
			copy(t, a)
			o = append(o, t)
			return
		}
		for _, w := range h[cur] {
			if _, ok := allSet[w]; !ok {
				continue
			}
			a = append(a, w)
			backtrack(w)
			a = a[:len(a)-1]
		}
	}
	backtrack(beginWord)
	return o
}
