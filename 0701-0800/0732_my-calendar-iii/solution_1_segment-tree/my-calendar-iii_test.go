package main

import "fmt"

func ExampleMyCalendarThree() {
	t := Constructor()
	fmt.Println(t.Book(10, 20))
	fmt.Println(t.Book(50, 60))
	fmt.Println(t.Book(10, 40))
	fmt.Println(t.Book(5, 15))
	fmt.Println(t.Book(5, 10))
	fmt.Println(t.Book(25, 55))
	// Output:
	// 1
	// 1
	// 2
	// 3
	// 3
	// 3
}
