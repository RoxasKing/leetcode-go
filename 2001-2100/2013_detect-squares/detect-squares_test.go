package main

import "fmt"

func ExampleDetectSquares() {
	ds := Constructor()
	ds.Add([]int{3, 10})
	ds.Add([]int{11, 2})
	ds.Add([]int{3, 2})
	fmt.Println(ds.Count([]int{11, 10}))
	fmt.Println(ds.Count([]int{14, 8}))
	ds.Add([]int{11, 2})
	fmt.Println(ds.Count([]int{11, 10}))
	// Output:
	// 1
	// 0
	// 2
}
