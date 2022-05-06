package main

import "fmt"

func ExampleRecentCounter() {
	rc := Constructor()
	fmt.Println(rc.Ping(1))
	fmt.Println(rc.Ping(100))
	fmt.Println(rc.Ping(3001))
	fmt.Println(rc.Ping(3002))
	// Output:
	// 1
	// 2
	// 3
	// 3
}
