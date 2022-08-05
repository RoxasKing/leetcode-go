package main

import "fmt"

func ExampleMyCircularQueue() {
	mcq := Constructor(3)
	fmt.Println(mcq.EnQueue(1))
	fmt.Println(mcq.EnQueue(2))
	fmt.Println(mcq.EnQueue(3))
	fmt.Println(mcq.EnQueue(4))
	fmt.Println(mcq.Rear())
	fmt.Println(mcq.IsFull())
	fmt.Println(mcq.DeQueue())
	fmt.Println(mcq.EnQueue(4))
	fmt.Println(mcq.Rear())
	// Output:
	// true
	// true
	// true
	// false
	// 3
	// true
	// true
	// true
	// 4
}
