package main

import "fmt"

func ExampleStockSpanner() {
	ss := NewStockSpanner()
	fmt.Println(ss.Next(100))
	fmt.Println(ss.Next(80))
	fmt.Println(ss.Next(60))
	fmt.Println(ss.Next(70))
	fmt.Println(ss.Next(60))
	fmt.Println(ss.Next(75))
	fmt.Println(ss.Next(85))
	// Output:
	// 1
	// 1
	// 1
	// 2
	// 1
	// 4
	// 6
}
