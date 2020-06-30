package leetcode

import "testing"

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("%s should in Trie\n", "apple")
		return
	}
	if trie.Search("app") {
		t.Errorf("%s shouldn't in Trie\n", "app")
		return
	}
	if !trie.StartsWith("app") {
		t.Errorf("prefix %s should have no word in Trie\n", "app")
		return
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("%s should in Trie\n", "app")
		return
	}
}
