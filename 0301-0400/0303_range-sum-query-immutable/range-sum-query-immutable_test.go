package main

import "fmt"

func ExampleNumArray() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
	// Output:
	// 1
	// -1
	// -3
}
