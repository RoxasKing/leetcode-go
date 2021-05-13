package main

import "fmt"

func ExampleTimeMap() {
	ts := Constructor()
	ts.Set("foo", "bar", 1)
	fmt.Println(ts.Get("foo", 1))
	fmt.Println(ts.Get("foo", 3))
	ts.Set("foo", "bar2", 4)
	fmt.Println(ts.Get("foo", 4))
	fmt.Println(ts.Get("foo", 5))
	// Output:
	// bar
	// bar
	// bar2
	// bar2
}
