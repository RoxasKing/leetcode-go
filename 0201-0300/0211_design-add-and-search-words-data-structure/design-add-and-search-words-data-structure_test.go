package main

import "fmt"

func ExampleWordDictionary() {
	wd := Constructor()
	wd.AddWord("at")
	wd.AddWord("and")
	wd.AddWord("an")
	wd.AddWord("add")
	fmt.Println(wd.Search("a"))
	fmt.Println(wd.Search(".at"))
	wd.AddWord("bat")
	fmt.Println(wd.Search(".at"))
	fmt.Println(wd.Search("an."))
	fmt.Println(wd.Search("a.d."))
	fmt.Println(wd.Search("b."))
	fmt.Println(wd.Search("a.d"))
	fmt.Println(wd.Search("."))
	// Output:
	// false
	// false
	// true
	// true
	// false
	// false
	// true
	// false
}
