package main

// Difficulty:
// Hard

// Tags:
// Trie

func palindromePairs(words []string) [][]int {
	trie := NewTrie()
	for i, word := range words {
		trie.Insert(word, i)
	}

	out := [][]int{}
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPalindrome(word[j:]) {
				if k := trie.Search(reverse(word[:j])); k != -1 && k != i {
					out = append(out, []int{i, k})
				}
			}
			if j > 0 && isPalindrome(word[:j]) {
				if k := trie.Search(reverse(word[j:])); k != -1 {
					out = append(out, []int{k, i})
				}
			}
		}
	}
	return out
}

func reverse(word string) string {
	chs := []byte(word)
	n := len(word)
	l, r := 0, n-1
	for l < r {
		chs[l], chs[r] = chs[r], chs[l]
		l++
		r--
	}
	return string(chs)
}

func isPalindrome(str string) bool {
	n := len(str)
	l, r := 0, n-1
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

type Trie struct {
	child [26]*Trie
	index int
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
		index: -1,
	}
}

func (t *Trie) Insert(word string, index int) {
	n := t
	for i := range word {
		idx := int(word[i] - 'a')
		if n.child[idx] == nil {
			n.child[idx] = NewTrie()
		}
		n = n.child[idx]
	}
	n.index = index
}

func (t *Trie) Search(word string) int {
	n := t
	for i := range word {
		idx := int(word[i] - 'a')
		if n.child[idx] == nil {
			return -1
		}
		n = n.child[idx]
	}
	return n.index
}
