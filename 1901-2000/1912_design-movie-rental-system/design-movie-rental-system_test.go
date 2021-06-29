package main

import "fmt"

func ExampleMovieRentingSystem() {
	c := Constructor(3, [][]int{{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5}})
	fmt.Println(c.Search(1))
	c.Rent(0, 1)
	c.Rent(1, 2)
	fmt.Println(c.Report())
	c.Drop(1, 2)
	fmt.Println(c.Search(2))
	// Output:
	// [1 0 2]
	// [[0 1] [1 2]]
	// [0 1]
}
