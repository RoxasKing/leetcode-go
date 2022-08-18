package main

import "fmt"

func ExampleMyCircularDeque() {
	mcd := Constructor(3)
	fmt.Println(mcd.InsertLast(1))
	fmt.Println(mcd.InsertLast(2))
	fmt.Println(mcd.InsertFront(3))
	fmt.Println(mcd.InsertFront(4))
	fmt.Println(mcd.GetRear())
	fmt.Println(mcd.IsFull())
	fmt.Println(mcd.DeleteLast())
	fmt.Println(mcd.InsertFront(4))
	fmt.Println(mcd.GetFront())
	// Output:
	// true
	// true
	// true
	// false
	// 2
	// true
	// true
	// true
	// 4
}
