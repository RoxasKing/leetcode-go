package main

import "fmt"

func ExampleSkiplist() {
	sl := Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(3)
	fmt.Println(sl.Search(0))
	sl.Add(4)
	fmt.Println(sl.Search(1))
	fmt.Println(sl.Erase(0))
	fmt.Println(sl.Erase(1))
	fmt.Println(sl.Search(1))
	// Output:
	// false
	// true
	// false
	// true
	// false
}
