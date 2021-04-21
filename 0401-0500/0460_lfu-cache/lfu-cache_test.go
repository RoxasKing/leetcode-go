package main

import "fmt"

func ExampleLFUCache() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	fmt.Println(lfu.Get(1))
	lfu.Put(3, 3)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 4)
	fmt.Println(lfu.Get(1))
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(4))
	// Output:
	// 1
	// -1
	// 3
	// -1
	// 3
	// 4
}
