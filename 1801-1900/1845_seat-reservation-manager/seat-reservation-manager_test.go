package main

import "fmt"

func ExampleSeatManager() {
	sm := Constructor(5)
	fmt.Println(sm.Reserve())
	fmt.Println(sm.Reserve())
	sm.Unreserve(2)
	fmt.Println(sm.Reserve())
	fmt.Println(sm.Reserve())
	fmt.Println(sm.Reserve())
	fmt.Println(sm.Reserve())
	sm.Unreserve(5)
	// Output:
	// 1
	// 2
	// 2
	// 3
	// 4
	// 5
}
