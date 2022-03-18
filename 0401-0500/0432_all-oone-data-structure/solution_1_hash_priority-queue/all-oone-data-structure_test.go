package main

import "fmt"

func ExampleAllOne() {
	ao := Constructor()
	ao.Inc("hello")
	ao.Inc("hello")
	fmt.Println(ao.GetMaxKey())
	fmt.Println(ao.GetMinKey())
	ao.Inc("leet")
	fmt.Println(ao.GetMaxKey())
	fmt.Println(ao.GetMinKey())

	fmt.Println()

	ao = Constructor()
	ao.Inc("a")
	ao.Inc("b")
	ao.Inc("b")
	ao.Inc("b")
	ao.Inc("b")
	ao.Dec("b")
	ao.Dec("b")
	fmt.Println(ao.GetMaxKey())
	fmt.Println(ao.GetMinKey())

	fmt.Println()

	ao = Constructor()
	ao.Inc("a")
	ao.Inc("b")
	ao.Inc("b")
	ao.Inc("c")
	ao.Inc("c")
	ao.Inc("c")
	ao.Dec("b")
	ao.Dec("b")
	fmt.Println(ao.GetMinKey())
	ao.Dec("a")
	fmt.Println(ao.GetMaxKey())
	fmt.Println(ao.GetMinKey())

	// Output:
	// hello
	// hello
	// hello
	// leet
	//
	// b
	// a
	//
	// a
	// c
	// c
}
