package main

/*
  Given a string band an array of smaller strings T, design a method to search b for each small string in T. Output positions of all strings in smalls that appear in big, where positions[i] is all positions of smalls[i].

  Example:
    Input:
      big = "mississippi"
      smalls = ["is","ppi","hi","sis","i","ssippi"]
    Output:  [[1,4],[8],[],[3],[1,4,7,10],[5]]

  Note:
    1. 0 <= len(big) <= 1000
    2. 0 <= len(smalls[i]) <= 1000
    3. The total number of characters in smalls will not exceed 100000.
    4. No duplicated strings in smalls.
    5. All characters are lowercase letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/multi-search-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
func multiSearch(big string, smalls []string) [][]int {
	t := NewTrie()
	for _, w := range smalls {
		t.Insert(w)
	}
	for i := range big {
		t.AddIdx(big[i:], i)
	}
	n := len(smalls)
	out := make([][]int, n)
	for i, w := range smalls {
		out[i] = t.Query(w)
	}
	return out
}

type Trie struct {
	child [26]*Trie
	isEnd bool
	idxs  []int
}

func NewTrie() *Trie {
	return &Trie{
		child: [26]*Trie{},
		isEnd: false,
		idxs:  []int{},
	}
}

func (t *Trie) Insert(word string) {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			t.child[idx] = NewTrie()
		}
		t = t.child[idx]
	}
	t.isEnd = true
}

func (t *Trie) AddIdx(str string, idx int) {
	for i := range str {
		j := str[i] - 'a'
		if t.child[j] == nil {
			break
		}
		t = t.child[j]
		if t.isEnd {
			t.idxs = append(t.idxs, idx)
		}
	}
}

func (t *Trie) Query(word string) []int {
	for i := range word {
		idx := word[i] - 'a'
		if t.child[idx] == nil {
			break
		}
		t = t.child[idx]
	}
	return t.idxs
}
