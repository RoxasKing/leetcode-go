package main

import "fmt"

func ExampleMKAverage() {
	mk := Constructor(3, 1)
	mk.AddElement(3)
	mk.AddElement(1)
	fmt.Println(mk.CalculateMKAverage())
	mk.AddElement(10)
	fmt.Println(mk.CalculateMKAverage())
	mk.AddElement(5)
	mk.AddElement(5)
	mk.AddElement(5)
	fmt.Println(mk.CalculateMKAverage())
	// Output:
	// -1
	// 3
	// 5
}
