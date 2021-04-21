package main

import "fmt"

func ExampleSolution() {
	nums := []int{1, 2, 3}
	s := Constructor(nums)
	s.Shuffle()
	fmt.Println(s.Reset())
	s.Shuffle()
	fmt.Println(s.Reset())
	// Output:
	// [1 2 3]
	// [1 2 3]
}
