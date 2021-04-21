package main

import "fmt"

func ExampleMajorityChecker() {
	mc := Constructor([]int{1, 1, 2, 2, 1, 1})
	fmt.Println(mc.Query(0, 5, 4))
	fmt.Println(mc.Query(0, 3, 3))
	fmt.Println(mc.Query(2, 3, 2))
	fmt.Println()

	mc = Constructor([]int{2, 1, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1, 2})
	fmt.Println(mc.Query(2, 9, 7))
	fmt.Println(mc.Query(9, 11, 2))
	fmt.Println(mc.Query(2, 11, 7))
	fmt.Println(mc.Query(3, 4, 2))
	fmt.Println(mc.Query(0, 1, 2))
	fmt.Println(mc.Query(6, 9, 3))
	fmt.Println(mc.Query(3, 12, 7))
	fmt.Println(mc.Query(3, 10, 6))
	fmt.Println(mc.Query(7, 11, 4))
	fmt.Println(mc.Query(0, 6, 4))

	// Output:
	// 1
	// -1
	// 2
	//
	// -1
	// 1
	// -1
	// -1
	// -1
	// 2
	// -1
	// -1
	// -1
	// 1
}
