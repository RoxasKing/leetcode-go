package leetcode

import (
	"fmt"
)

func ExampleLRUCache1() {
	obj := NewLRUCache(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	// Output:
	// 1
	// -1
	// -1
}

func ExampleLRUCache2() {
	obj := NewLRUCache(1)
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Put(3, 2)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	// Output:
	// 1
	// -1
	// 2
}
