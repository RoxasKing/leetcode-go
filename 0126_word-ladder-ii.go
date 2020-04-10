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
	wordDict := make(map[string]struct{})
	for _, w := range wordList {
		wordDict[w] = struct{}{}
	}
	if _, ok := wordDict[endWord]; !ok {
		return nil
	}
	delete(wordDict, endWord)
	src, dst := map[string]struct{}{beginWord: {}}, map[string]struct{}{endWord: {}}
	dict := make(map[string][]string)
	var isEnd, reverse bool
	for len(src) != 0 && len(dst) != 0 && !isEnd {
		if len(src) > len(dst) {
			src, dst = dst, src
			reverse = !reverse
		}
		for w := range src {
			delete(wordDict, w)
		}
		newSrc := make(map[string]struct{})
		for w := range src {
			bytes := []byte(w)
			for i := range bytes {
				for j := 'a'; j <= 'z'; j++ {
					bytes[i] = byte(j)
					target := string(bytes)
					if _, ok := dst[target]; ok {
						if reverse {
							dict[target] = append(dict[target], w)
						} else {
							dict[w] = append(dict[w], target)
						}
						isEnd = true
						continue
					}
					if _, ok := wordDict[target]; ok {
						newSrc[target] = struct{}{}
						if reverse {
							dict[target] = append(dict[target], w)
						} else {
							dict[w] = append(dict[w], target)
						}
					}
				}
				bytes[i] = w[i]
			}
		}
		src = newSrc
	}
	var out [][]string
	cur := []string{beginWord}
	var recur func([]string)
	recur = func(nexts []string) {
		if len(nexts) == 0 {
			return
		}
		for _, next := range nexts {
			cur = append(cur, next)
			if next == endWord {
				tmp := make([]string, len(cur))
				copy(tmp, cur)
				out = append(out, tmp)
			}
			recur(dict[next])
			cur = cur[:len(cur)-1]
		}
	}
	recur(dict[beginWord])
	return out
}
