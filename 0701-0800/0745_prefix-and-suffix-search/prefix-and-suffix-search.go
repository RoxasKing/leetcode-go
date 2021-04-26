package main

/*
  Design a special dictionary which has some words and allows you to search the words in it by a prefix and a suffix.

  Implement the WordFilter class:
    1. WordFilter(string[] words) Initializes the object with the words in the dictionary.
    2. f(string prefix, string suffix) Returns the index of the word in the dictionary which has the prefix prefix and the suffix suffix. If there is more than one valid index, return the largest of them. If there is no such word in the dictionary, return -1.

  Example 1:
    Input
      ["WordFilter", "f"]
      [[["apple"]], ["a", "e"]]
    Output
      [null, 0]
    Explanation
      WordFilter wordFilter = new WordFilter(["apple"]);
      wordFilter.f("a", "e"); // return 0, because the word at index 0 has prefix = "a" and suffix = 'e".

  Constraints:
    1. 1 <= words.length <= 15000
    2. 1 <= words[i].length <= 10
    3. 1 <= prefix.length, suffix.length <= 10
    4. words[i], prefix and suffix consist of lower-case English letters only.
    5. At most 15000 calls will be made to the function f.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/prefix-and-suffix-search
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
type WordFilter struct {
	trie *Trie
}

func Constructor(words []string) WordFilter {
	t := &Trie{}
	for i, word := range words {
		for j := range word {
			t.Insert(word[j:]+"#"+word, i)
		}
	}
	return WordFilter{
		trie: t,
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	return this.trie.Search(suffix + "#" + prefix)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

type Trie struct {
	child  [27]*Trie
	isEnd  bool
	weight int
}

func (t *Trie) Insert(word string, weight int) {
	for i := range word {
		idx := 0
		if word[i] != '#' {
			idx = int(word[i] - 'a' + 1)
		}
		if t.child[idx] == nil {
			t.child[idx] = &Trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
	t.weight = weight
}

func (t *Trie) Search(prefix string) int {
	for i := range prefix {
		idx := 0
		if prefix[i] != '#' {
			idx = int(prefix[i] - 'a' + 1)
		}
		if t.child[idx] == nil {
			return -1
		}
		t = t.child[idx]
	}

	out := -1
	q := []*Trie{t}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		if t.isEnd {
			out = Max(out, t.weight)
		}
		for i := 1; i <= 26; i++ {
			if t.child[i] != nil {
				q = append(q, t.child[i])
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
