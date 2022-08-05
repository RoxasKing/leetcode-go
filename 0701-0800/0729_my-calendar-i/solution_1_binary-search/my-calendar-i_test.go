package main

import "fmt"

func ExampleMyCalendar() {
	mc := Constructor()
	fmt.Println(mc.Book(10, 20))
	fmt.Println(mc.Book(15, 25))
	fmt.Println(mc.Book(20, 30))
	// Output:
	// true
	// false
	// true
}
