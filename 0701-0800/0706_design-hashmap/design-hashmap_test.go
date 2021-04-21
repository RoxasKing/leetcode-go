package main

import "fmt"

func ExampleMyHashMap() {
	m := Constructor()
	m.Put(1, 1)
	m.Put(2, 2)
	fmt.Println(m.Get(1))
	fmt.Println(m.Get(3))
	m.Put(2, 1)
	fmt.Println(m.Get(2))
	m.Remove(2)
	fmt.Println(m.Get(2))
	// Output:
	// 1
	// -1
	// 1
	// -1
}
