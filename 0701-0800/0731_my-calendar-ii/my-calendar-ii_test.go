package main

import "fmt"

func ExampleMyCalendarTwo() {
	mct := Constructor()
	fmt.Println(mct.Book(10, 20))
	fmt.Println(mct.Book(50, 60))
	fmt.Println(mct.Book(10, 40))
	fmt.Println(mct.Book(5, 15))
	fmt.Println(mct.Book(5, 10))
	fmt.Println(mct.Book(25, 55))
	// Output:
	// true
	// true
	// true
	// false
	// true
	// true
}
