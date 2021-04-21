package main

import "fmt"

func ExampleKthLargest() {
	kl := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kl.Add(3))
	fmt.Println(kl.Add(5))
	fmt.Println(kl.Add(10))
	fmt.Println(kl.Add(9))
	fmt.Println(kl.Add(4))
	// Output:
	// 4
	// 5
	// 5
	// 8
	// 8
}
