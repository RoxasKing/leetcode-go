package leetcode

/*
  给定一个单词列表，我们将这个列表编码成一个索引字符串 S 与一个索引列表 A。
  例如，如果这个列表是 ["time", "me", "bell"]，我们就可以将其表示为 S = "time#bell#" 和 indexes = [0, 2, 5]。
  对于每一个索引，我们可以通过从字符串 S 中索引的位置开始读取字符串，直到 "#" 结束，来恢复我们之前的单词列表。
  那么成功对给定单词列表进行编码的最小字符串长度是多少呢？
*/

func minimumLengthEncoding(words []string) int {
	dict := make(map[string]struct{})
	for _, word := range words {
		dict[word] = struct{}{}
	}
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			delete(dict, word[i:])
		}
	}
	var length int
	for k := range dict {
		length += len(k) + 1
	}
	return length
}

// Trie
func minimumLengthEncoding2(words []string) int {
	trie := newTrieNode()
	nodes := make(map[*TrieNode]int)
	for i, word := range words {
		cur := trie
		for j := len(word) - 1; j >= 0; j-- {
			cur = cur.get(word[j])
		}
		nodes[cur] = i
	}

	var length int
	for node := range nodes {
		if node.count == 0 {
			length += len(words[nodes[node]]) + 1
		}
	}
	return length
}

// TrieNode ...
type TrieNode struct {
	children []*TrieNode
	count    int
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make([]*TrieNode, 26),
		count:    0,
	}
}

func (t *TrieNode) get(c byte) *TrieNode {
	if t.children[c-'a'] == nil {
		t.children[c-'a'] = newTrieNode()
		t.count++
	}
	return t.children[c-'a']
}
