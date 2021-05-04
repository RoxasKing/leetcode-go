package main

import "fmt"

func ExampleMyStack() {
	ms := Constructor()
	ms.Push(1)
	ms.Push(2)
	fmt.Println(ms.Top())
	fmt.Println(ms.Pop())
	fmt.Println(ms.Empty())
	// Output:
	// 2
	// 2
	// false
}
