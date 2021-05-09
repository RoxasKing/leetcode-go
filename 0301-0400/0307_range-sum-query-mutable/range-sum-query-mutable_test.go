package main

import "fmt"

func ExampleNumArray() {
	na := Constructor([]int{1, 3, 5})
	fmt.Println(na.SumRange(0, 2))
	na.Update(1, 2)
	fmt.Println(na.SumRange(0, 2))

	fmt.Println()
	na = Constructor([]int{9, -8})
	na.Update(0, 3)
	fmt.Println(na.SumRange(1, 1))
	fmt.Println(na.SumRange(0, 1))
	na.Update(1, -3)
	fmt.Println(na.SumRange(0, 1))

	// Output:
	// 9
	// 8
	//
	// -8
	// -5
	// 0
}
