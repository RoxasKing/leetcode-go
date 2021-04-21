package main

/*
  A transformation sequence from word beginWord to word endWord using a dictionary wordList is a sequence of words beginWord -> s1 -> s2 -> ... -> sk such that:
    1. Every adjacent pair of words differs by a single letter.
    2. Every si for 1 <= i <= k is in wordList. Note that beginWord does not need to be in wordList.
    3. sk == endWord
  Given two words, beginWord and endWord, and a dictionary wordList, return all the shortest transformation sequences from beginWord to endWord, or an empty list if no such sequence exists. Each sequence should be returned as a list of the words [beginWord, s1, s2, ..., sk].

  Example 1:
    Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
    Output: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
    Explanation: There are 2 shortest transformation sequences:
      "hit" -> "hot" -> "dot" -> "dog" -> "cog"
      "hit" -> "hot" -> "lot" -> "log" -> "cog"

  Example 2:
    Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
    Output: []
    Explanation: The endWord "cog" is not in wordList, therefore there is no valid transformation sequence.

  Constraints:
    1. 1 <= beginWord.length <= 10
    2. endWord.length == beginWord.length
    3. 1 <= wordList.length <= 5000
    4. wordList[i].length == beginWord.length
    5. beginWord, endWord, and wordList[i] consist of lowercase English letters.
    6. beginWord != endWord
    7. All the words in wordList are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/word-ladder-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Backtracking
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
