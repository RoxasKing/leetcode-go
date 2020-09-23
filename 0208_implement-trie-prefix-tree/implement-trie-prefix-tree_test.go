package main

import "fmt"

func ExampleTrie() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	// Output:
	// true
	// false
	// true
	// true
	// true
}
