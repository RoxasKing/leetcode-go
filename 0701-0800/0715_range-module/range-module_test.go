package main

import "fmt"

func ExampleRangeModule() {
	rm := Constructor()
	rm.RemoveRange(4, 8)
	rm.AddRange(1, 10)
	fmt.Println(rm.QueryRange(1, 7))
	rm.AddRange(2, 3)
	rm.RemoveRange(2, 3)
	fmt.Println(rm.QueryRange(3, 7))
	fmt.Println(rm.QueryRange(8, 9))
	fmt.Println(rm.QueryRange(7, 9))
	fmt.Println(rm.QueryRange(6, 9))
	rm.AddRange(2, 3)
	rm.RemoveRange(1, 8)
	// Output:
	// true
	// true
	// true
	// true
	// true
}
