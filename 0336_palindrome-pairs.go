package leetcode

/*
  给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词， words[i] + words[j] ，可拼接成回文串。
*/

// Trie
func palindromePairs(words []string) [][]int {
	type Trie struct {
		next  [26]*Trie
		index int
	}
	trie := &Trie{next: [26]*Trie{}, index: -1} // build trie
	for i, word := range words {
		ptr := trie
		for j := 0; j < len(word); j++ {
			w := int(word[j] - 'a')
			if ptr.next[w] == nil {
				ptr.next[w] = &Trie{next: [26]*Trie{}, index: -1}
			}
			ptr = ptr.next[w]
		}
		ptr.index = i
	}

	searchIndex := func(s string) int { // find in trie
		ptr := trie
		for i := len(s) - 1; i >= 0; i-- {
			w := int(s[i] - 'a')
			if ptr.next[w] == nil {
				return -1
			}
			ptr = ptr.next[w]
		}
		return ptr.index
	}

	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)>>1; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}

	var out [][]int
	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			if isPalindrome(word[j:]) {
				if index := searchIndex(word[:j]); index != -1 && index != i {
					out = append(out, []int{i, index})
				}
			}
			if j != 0 && isPalindrome(word[:j]) {
				if index := searchIndex(word[j:]); index != -1 && index != i {
					out = append(out, []int{index, i})
				}
			}
		}
	}
	return out
}
