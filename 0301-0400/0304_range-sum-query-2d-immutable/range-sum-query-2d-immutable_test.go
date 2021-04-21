package main

import "fmt"

func ExampleNumMatrix() {
	nm := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(nm.SumRegion(2, 1, 4, 3))
	fmt.Println(nm.SumRegion(1, 1, 2, 2))
	fmt.Println(nm.SumRegion(1, 2, 2, 4))
	fmt.Println()
	nm = Constructor([][]int{
		{-4, -5},
	})
	fmt.Println(nm.SumRegion(0, 0, 0, 0))
	fmt.Println(nm.SumRegion(0, 0, 0, 1))
	fmt.Println(nm.SumRegion(0, 1, 0, 1))
	fmt.Println()
	// Output:
	// 8
	// 11
	// 12
	//
	// -4
	// -9
	// -5
}
