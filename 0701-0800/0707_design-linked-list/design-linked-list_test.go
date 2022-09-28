package main

import "fmt"

func ExampleMyLinkedList() {
	mll := Constructor()
	mll.AddAtHead(1)
	mll.AddAtTail(3)
	mll.AddAtIndex(1, 2)
	fmt.Println(mll.Get(1))
	mll.DeleteAtIndex(1)
	fmt.Println(mll.Get(1))
	// Output:
	// 2
	// 3
}
