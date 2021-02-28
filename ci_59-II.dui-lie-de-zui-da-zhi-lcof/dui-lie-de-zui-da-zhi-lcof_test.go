package main

import "fmt"

func ExampleMaxQueue() {
	mq := Constructor()
	mq.Push_back(1)
	mq.Push_back(2)
	fmt.Println(mq.Max_value())
	fmt.Println(mq.Pop_front())
	fmt.Println(mq.Max_value())
	// Output:
	// 2
	// 1
	// 2
}
