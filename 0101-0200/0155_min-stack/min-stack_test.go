package main

import (
	"fmt"
)

func ExampleMinStack() {
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	fmt.Println(minstack.GetMin())
	minstack.Pop()
	fmt.Println(minstack.Top())
	fmt.Println(minstack.GetMin())
	minstack.Pop()
	fmt.Println(minstack.Top())
	fmt.Println(minstack.GetMin())
	// Output:
	// -3
	// 0
	// -2
	// -2
	// -2
}
