package leetcode

/*
  给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
    每次转换只能改变一个字母。
    转换过程中的中间单词必须是字典中的单词。
  说明:
    如果不存在这样的转换序列，返回一个空列表。
    所有单词具有相同的长度。
    所有单词只由小写字母组成。
    字典中不存在重复的单词。
    你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return nil
	}
	delete(wordSet, endWord)
	src, dst := map[string]bool{beginWord: true}, map[string]bool{endWord: true}
	dict := make(map[string][]string)
	var reverse, isEnd bool
	for len(src) != 0 && len(dst) != 0 && !isEnd {
		if len(src) > len(dst) {
			src, dst = dst, src
			reverse = !reverse
		}
		for word := range src {
			delete(wordSet, word)
		}
		newSrc := make(map[string]bool)
		for word := range src {
			bytes := []byte(word)
			for i := range bytes {
				for j := 'a'; j <= 'z'; j++ {
					bytes[i] = byte(j)
					targetWord := string(bytes)
					if dst[targetWord] {
						if reverse {
							dict[targetWord] = append(dict[targetWord], word)
						} else {
							dict[word] = append(dict[word], targetWord)
						}
						isEnd = true
					} else if wordSet[targetWord] {
						newSrc[targetWord] = true
						if reverse {
							dict[targetWord] = append(dict[targetWord], word)
						} else {
							dict[word] = append(dict[word], targetWord)
						}
					}
				}
				bytes[i] = word[i]
			}
		}
		src = newSrc
	}
	var out [][]string
	cur := []string{beginWord}
	var recur func([]string)
	recur = func(words []string) {
		if len(words) == 0 {
			return
		}
		for _, word := range words {
			cur = append(cur, word)
			if word == endWord {
				tmp := make([]string, len(cur))
				copy(tmp, cur)
				out = append(out, tmp)
			} else {
				recur(dict[word])
			}
			cur = cur[:len(cur)-1]
		}
	}
	recur(dict[beginWord])
	return out
}
