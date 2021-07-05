package main

// Tags:
// Backtracking
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wSet := make(map[string]bool)
	for _, word := range wordList {
		wSet[word] = true
	}
	if !wSet[endWord] {
		return nil
	}
	delete(wSet, endWord)

	src, dst := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	dict := make(map[string][]string)
	var rev, ok bool

	for len(src) != 0 && len(dst) != 0 && !ok {
		if len(src) > len(dst) {
			src, dst = dst, src
			rev = !rev
		}
		for word := range src {
			delete(wSet, word)
		}
		nSrc := make(map[string]bool)
		for word := range src {
			chs := []byte(word)
			for i := range chs {
				for ch := byte('a'); ch <= 'z'; ch++ {
					chs[i] = ch
					next := string(chs)
					if dst[next] {
						if rev {
							dict[next] = append(dict[next], word)
						} else {
							dict[word] = append(dict[word], next)
						}
						ok = true
					} else if wSet[next] {
						nSrc[next] = true
						if rev {
							dict[next] = append(dict[next], word)
						} else {
							dict[word] = append(dict[word], next)
						}
					}
				}
				chs[i] = word[i]
			}
		}
		src = nSrc
	}

	out := [][]string{}
	cur := []string{beginWord}
	dfs(dict, &out, endWord, cur, dict[beginWord])
	return out
}

func dfs(dict map[string][]string, out *[][]string, endWord string, cur []string, words []string) {
	if len(words) == 0 {
		return
	}
	for _, word := range words {
		cur = append(cur, word)
		if word == endWord {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			*out = append(*out, tmp)
		} else {
			dfs(dict, out, endWord, cur, dict[word])
		}
		cur = cur[:len(cur)-1]
	}
}
