package main

// Tags:
// Graph
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordDict := make(map[string]bool)
	for _, word := range wordList {
		wordDict[word] = true
	}
	if !wordDict[endWord] {
		return 0
	}
	delete(wordDict, endWord)
	src := map[string]bool{beginWord: true}
	dst := map[string]bool{endWord: true}
	out := 1
	for len(src) != 0 && len(dst) != 0 {
		out++
		if len(src) > len(dst) {
			src, dst = dst, src
		}
		for word := range src {
			delete(wordDict, word)
		}
		newSrc := map[string]bool{}
		for word := range src {
			letters := []byte(word)
			for i := range letters {
				for j := 0; j < 26; j++ {
					letters[i] = byte(j) + 'a'
					next := string(letters)
					if dst[next] {
						return out
					}
					if wordDict[next] {
						newSrc[next] = true
					}
				}
				letters[i] = word[i]
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
		*queue = (*queue)[1:]
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
