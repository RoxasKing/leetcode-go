package main

import "fmt"

func ExampleMagicDictionary() {
	md := Constructor()
	md.BuildDict([]string{"hello", "leetcode"})
	fmt.Println(md.Search("hello"))
	fmt.Println(md.Search("hhllo"))
	fmt.Println(md.Search("hell"))
	fmt.Println(md.Search("leetcoded"))
	// Output:
	// false
	// true
	// false
	// false
}
