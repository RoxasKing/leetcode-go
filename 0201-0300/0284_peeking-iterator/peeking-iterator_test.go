package main

import "fmt"

func ExamplePeekingIterator() {
	iter := &Iterator{arr: []int{1, 2, 3}}
	pi := Constructor(iter)
	fmt.Println(pi.next())
	fmt.Println(pi.peek())
	fmt.Println(pi.next())
	fmt.Println(pi.next())
	fmt.Println(pi.hasNext())
	// Output:
	// 1
	// 2
	// 2
	// 3
	// false
}
