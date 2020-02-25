package leetcode

/*
  给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
    每次转换只能改变一个字母。
    转换过程中的中间单词必须是字典中的单词。
  说明:
    如果不存在这样的转换序列，返回 0。
    所有单词具有相同的长度。
    所有单词只由小写字母组成。
    字典中不存在重复的单词。
    你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordDict := make(map[string]struct{})
	for _, w := range wordList {
		wordDict[w] = struct{}{}
	}
	if _, ok := wordDict[endWord]; !ok {
		return 0
	}
	src := map[string]struct{}{beginWord: struct{}{}}
	dst := map[string]struct{}{endWord: struct{}{}}
	out := 1
	for len(src) != 0 && len(dst) != 0 {
		out++
		if len(src) > len(dst) {
			src, dst = dst, src
		}
		for w := range src {
			delete(wordDict, w)
		}
		newSrc := make(map[string]struct{})
		for w := range src {
			bytes := []byte(w)
			for i := 0; i < len(bytes); i++ {
				for j := 0; j < 26; j++ {
					bytes[i] = byte(j + 'a')
					target := string(bytes)
					if _, ok := dst[target]; ok {
						return out
					}
					if _, ok := wordDict[target]; ok {
						newSrc[target] = struct{}{}
					}
				}
				bytes[i] = w[i]
			}
		}
		src = newSrc
	}
	return 0
}

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	var flag bool
	for _, word := range wordList {
		if word == endWord {
			flag = true
		}
	}
	if !flag {
		return 0
	}
	allcomboDict := make(map[string][]string)
	for _, word := range wordList {
		for i := range beginWord {
			newWord := word[:i] + "*" + word[i+1:]
			allcomboDict[newWord] = append(allcomboDict[newWord], word)
		}
	}
	type node struct {
		word  string
		level int
	}
	visitWordNode := func(queue *[]node, visited, visitedOther *map[string]int) int {
		word, level := (*queue)[0].word, (*queue)[0].level
		*queue = (*queue)[1:]
		for i := range beginWord {
			newWord := word[:i] + "*" + word[i+1:]
			for _, adjacentWord := range allcomboDict[newWord] {
				if other, ok := (*visitedOther)[adjacentWord]; ok {
					return level + other
				}
				if _, ok := (*visited)[adjacentWord]; !ok {
					(*visited)[adjacentWord] = level + 1
					(*queue) = append((*queue), node{adjacentWord, level + 1})
				}
			}
		}
		return -1
	}
	queueBegin := []node{{beginWord, 1}}
	queueEnd := []node{{endWord, 1}}
	visitedBegin := map[string]int{beginWord: 1}
	visitedEnd := map[string]int{endWord: 1}
	for len(queueBegin) != 0 && len(queueEnd) != 0 {
		out := visitWordNode(&queueBegin, &visitedBegin, &visitedEnd)
		if out > -1 {
			return out
		}
		out = visitWordNode(&queueEnd, &visitedEnd, &visitedBegin)
		if out > -1 {
			return out
		}
	}
	return 0
}

func ladderLength3(beginWord string, endWord string, wordList []string) int {
	allComboDict := make(map[string][]string)
	for _, word := range wordList {
		for i := range beginWord {
			newWord := word[:i] + "*" + word[i+1:]
			allComboDict[newWord] = append(allComboDict[newWord], word)
		}
	}
	type node struct {
		word  string
		level int
	}
	queue := []node{{beginWord, 1}}
	visited := map[string]bool{beginWord: true}
	for len(queue) != 0 {
		word, level := queue[0].word, queue[0].level
		queue = queue[1:]
		for i := range beginWord {
			newWord := word[:i] + "*" + word[i+1:]
			for _, adjacentWord := range allComboDict[newWord] {
				if adjacentWord == endWord {
					return level + 1
				}
				if _, ok := visited[adjacentWord]; !ok {
					visited[adjacentWord] = true
					queue = append(queue, node{adjacentWord, level + 1})
				}
			}
		}
	}
	return 0
}
