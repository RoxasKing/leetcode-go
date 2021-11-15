package main

import "fmt"

func ExampleCombinationIterator() {
	ci := Constructor("abc", 2)
	fmt.Println(ci.Next())
	fmt.Println(ci.HasNext())
	fmt.Println(ci.Next())
	fmt.Println(ci.HasNext())
	fmt.Println(ci.Next())
	fmt.Println(ci.HasNext())
	// Output:
	// ab
	// true
	// ac
	// true
	// bc
	// false
}
