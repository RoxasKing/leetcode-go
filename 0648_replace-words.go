package main

import (
	"strings"
)

/*
  在英语中，我们有一个叫做 词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。

  现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。

  你需要输出替换之后的句子。

  提示：
    输入只包含小写字母。
    1 <= dict.length <= 1000
    1 <= dict[i].length <= 100
    1 <= 句中词语数 <= 1000
    1 <= 句中词语长度 <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/replace-words
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
func replaceWords(dict []string, sentence string) string {
	type Trie struct {
		branch [26]*Trie
		word   string
	}
	trie := new(Trie)
	for _, root := range dict {
		cur := trie
		for i := range root {
			if cur.branch[root[i]-'a'] == nil {
				cur.branch[root[i]-'a'] = new(Trie)
			}
			cur = cur.branch[root[i]-'a']
		}
		cur.word = root
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		cur := trie
		for i := range word {
			if cur.branch[word[i]-'a'] == nil || cur.word != "" {
				break
			}
			cur = cur.branch[word[i]-'a']
		}
		if cur.word != "" {
			words[i] = cur.word
		}
	}
	return strings.Join(words, " ")
}

// Prefix Hash
func replaceWords2(dict []string, sentence string) string {
	hasPrefix := make(map[string]bool, len(dict))
	hasLength := make([]bool, 101)
	for i := range dict {
		hasPrefix[dict[i]] = true
		hasLength[len(dict[i])] = true
	}
	lengths := make([]int, 101)
	var index int
	for i, ok := range hasLength {
		if ok {
			lengths[index] = i
			index++
		}
	}
	lengths = lengths[:index]

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 0; j < len(lengths) && lengths[j] < len(word); j++ {
			if hasPrefix[word[:lengths[j]]] {
				words[i] = word[:lengths[j]]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
