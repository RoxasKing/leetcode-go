package main

import "fmt"

func ExampleFindSumPairs() {
	fsp := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(fsp.Count(7))
	fsp.Add(3, 2)
	fmt.Println(fsp.Count(8))
	fmt.Println(fsp.Count(4))
	fsp.Add(0, 1)
	fsp.Add(1, 1)
	fmt.Println(fsp.Count(7))
	// Output:
	// 8
	// 2
	// 1
	// 11
}
