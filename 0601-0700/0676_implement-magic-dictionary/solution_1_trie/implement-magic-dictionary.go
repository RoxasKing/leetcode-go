package main

/*
  Design a data structure that is initialized with a list of different words. Provided a string, you should determine if you can change exactly one character in this string to match any word in the data structure.

  Implement the MagicDictionary class:
    1. MagicDictionary() Initializes the object.
    2. void buildDict(String[] dictionary) Sets the data structure with an array of distinct strings dictionary.
    3. bool search(String searchWord) Returns true if you can change exactly one character in searchWord to match any string in the data structure, otherwise returns false.


  Example 1:
    Input
      ["MagicDictionary", "buildDict", "search", "search", "search", "search"]
      [[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
    Output
      [null, null, false, true, false, false]
    Explanation
      MagicDictionary magicDictionary = new MagicDictionary();
      magicDictionary.buildDict(["hello", "leetcode"]);
      magicDictionary.search("hello"); // return False
      magicDictionary.search("hhllo"); // We can change the second 'h' to 'e' to match "hello" so we return True
      magicDictionary.search("hell"); // return False
      magicDictionary.search("leetcoded"); // return False

  Constraints:
    1. 1 <= dictionary.length <= 100
    2. 1 <= dictionary[i].length <= 100
    3. dictionary[i] consists of only lower-case English letters.
    4. All the strings in dictionary are distinct.
    5. 1 <= searchWord.length <= 100
    6. searchWord consists of only lower-case English letters.
    7. buildDict will be called only once before search.
    8. At most 100 calls will be made to search.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/implement-magic-dictionary
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Trie
type MagicDictionary struct {
	t *trie
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{t: &trie{}}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, w := range dictionary {
		this.t.insert(w)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	q := []*triePair{{t: this.t, changed: false}}
	for i := range searchWord {
		idx := int(searchWord[i] - 'a')
		size := len(q)
		for _, p := range q {
			if p.t.child[idx] != nil {
				q = append(q, &triePair{t: p.t.child[idx], changed: p.changed})
			}
			if !p.changed {
				for i := 0; i < 26; i++ {
					if i == idx || p.t.child[i] == nil {
						continue
					}
					q = append(q, &triePair{t: p.t.child[i], changed: true})
				}
			}
		}
		q = q[size:]
	}

	for _, p := range q {
		if p.t.isEnd && p.changed {
			return true
		}
	}
	return false
}

type triePair struct {
	t       *trie
	changed bool
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */

type trie struct {
	child [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	for i := range word {
		idx := int(word[i] - 'a')
		if t.child[idx] == nil {
			t.child[idx] = &trie{}
		}
		t = t.child[idx]
	}
	t.isEnd = true
}
