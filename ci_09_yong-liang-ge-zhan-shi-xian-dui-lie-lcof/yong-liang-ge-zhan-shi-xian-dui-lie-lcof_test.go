package main

import (
	"fmt"
)

func ExampleCQueue() {
	cq := Constructor()
	cq.AppendTail(3)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
	// Output:
	// 3
	// -1
}
