package main

import "sort"

/*
  A valid encoding of an array of words is any reference string s and array of indices indices such that:
    1. words.length == indices.length
    2. The reference string s ends with the '#' character.
    3. For each index indices[i], the substring of s starting from indices[i] and up to (but not including) the next '#' character is equal to words[i].
  Given an array of words, return the length of the shortest reference string s possible of any valid encoding of words.

  Example 1:
    Input: words = ["time", "me", "bell"]
    Output: 10
    Explanation: A valid encoding would be s = "time#bell#" and indices = [0, 2, 5].
      words[0] = "time", the substring of s starting from indices[0] = 0 to the next '#' is underlined in "time#bell#"
      words[1] = "me", the substring of s starting from indices[1] = 2 to the next '#' is underlined in "time#bell#"
      words[2] = "bell", the substring of s starting from indices[2] = 5 to the next '#' is underlined in "time#bell#"

  Example 2:
    Input: words = ["t"]
    Output: 2
    Explanation: A valid encoding would be s = "t#" and indices = [0].

  Constraints:
    1. 1 <= words.length <= 2000
    2. 1 <= words[i].length <= 7
    3. words[i] consists of only lowercase letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/short-encoding-of-words
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash
func minimumLengthEncoding(words []string) int {
	mark := make(map[string]struct{})
	for _, word := range words {
		mark[word] = struct{}{}
	}
	for _, word := range words {
		for word != "" {
			word = word[1:]
			delete(mark, word)
		}
	}
	out := 0
	for w := range mark {
		out += len(w) + 1
	}
	return out
}

// Trie
func minimumLengthEncoding2(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	out := 0
	visited := make(map[string]bool)
	for _, word := range words {
		if visited[word] {
			continue
		}
		visited[word] = true
		if trie.TouchEnd(word) {
			out += len(word) + 1
		}
	}
	return out
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
	}
}

func (trie *Trie) Insert(word string) {
	t := trie
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = NewTrie()
		}
		t = t.child[idx]
		t.isEnd = false
	}
	t.isEnd = true
}

func (trie *Trie) TouchEnd(word string) bool {
	t := trie
	n := len(word)
	for i := n - 1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			break
		}
		t = t.child[idx]
	}
	return t.isEnd
}
