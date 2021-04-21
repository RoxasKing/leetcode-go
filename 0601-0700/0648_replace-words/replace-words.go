package main

import (
	"strings"
)

/*
  In English, we have a concept called root, which can be followed by some other word to form another longer word - let's call this word successor. For example, when the root "an" is followed by the successor word "other", we can form a new word "another".

  Given a dictionary consisting of many roots and a sentence consisting of words separated by spaces, replace all the successors in the sentence with the root forming it. If a successor can be replaced by more than one root, replace it with the root that has the shortest length.

  Return the sentence after the replacement.

  Example 1:
    Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
    Output: "the cat was rat by the bat"

  Example 2:
    Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
    Output: "a a b c"

  Example 3:
    Input: dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
    Output: "a a a a a a a a bbb baba a"

  Example 4:
    Input: dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
    Output: "the cat was rat by the bat"

  Example 5:
    Input: dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
    Output: "it is ab that this solution is ac"

  Constraints:
    1. 1 <= dictionary.length <= 1000
    2. 1 <= dictionary[i].length <= 100
    3. dictionary[i] consists of only lower-case letters.
    4. 1 <= sentence.length <= 10^6
    5. sentence consists of only lower-case letters and spaces.
    6. The number of words in sentence is in the range [1, 1000]
    7. The length of each word in sentence is in the range [1, 1000]
    8. Each two consecutive words in sentence will be separated by exactly one space.
    9. sentence does not have leading or trailing spaces.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/replace-words
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
func replaceWords(dict []string, sentence string) string {
	trie := NewTrie()
	for _, w := range dict {
		trie.Insert(w)
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if res := trie.Search(word); res != "" {
			words[i] = res
		}
	}
	return strings.Join(words, " ")
}

type Trie struct {
	child  [26]*Trie
	prefix string
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	trie := t
	for i := range word {
		idx := int(word[i] - 'a')
		if trie.child[idx] == nil {
			trie.child[idx] = NewTrie()
		}
		trie = trie.child[idx]
	}
	trie.prefix = word
}

func (t *Trie) Search(word string) string {
	trie := t
	for i := range word {
		idx := int(word[i] - 'a')
		if trie.child[idx] == nil {
			break
		}
		trie = trie.child[idx]
		if trie.prefix != "" {
			break
		}
	}
	return trie.prefix
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
