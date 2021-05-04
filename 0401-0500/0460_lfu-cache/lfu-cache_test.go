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

	fmt.Println()
	lfu = Constructor(1)
	lfu.Put(2, 1)
	fmt.Println(lfu.Get(2))
	lfu.Put(3, 2)
	fmt.Println(lfu.Get(2))
	fmt.Println(lfu.Get(3))

	fmt.Println()
	lfu = Constructor(2)
	lfu.Put(2, 1)
	lfu.Put(3, 2)
	fmt.Println(lfu.Get(3))
	fmt.Println(lfu.Get(3))
	lfu.Put(4, 3)
	fmt.Println(lfu.Get(4))
	fmt.Println(lfu.Get(4))
	fmt.Println(lfu.Get(4))
	lfu.Put(5, 4)
	fmt.Println(lfu.Get(3))

	// Output:
	// 1
	// -1
	// 3
	// -1
	// 3
	// 4
	//
	// 1
	// -1
	// 2
	//
	// 2
	// 2
	// 3
	// 3
	// 3
	// -1
}
