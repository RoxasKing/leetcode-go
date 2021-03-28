package main

import "fmt"

func ExampleMyHashSet() {
	mh := Constructor()
	mh.Add(1)
	mh.Add(2)
	fmt.Println(mh.Contains(1))
	fmt.Println(mh.Contains(3))
	mh.Add(2)
	fmt.Println(mh.Contains(2))
	mh.Remove(2)
	fmt.Println(mh.Contains(2))
	// Output:
	// true
	// false
	// true
	// false
}
