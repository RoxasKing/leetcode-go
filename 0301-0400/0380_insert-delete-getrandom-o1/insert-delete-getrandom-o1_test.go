package main

import "fmt"

func ExampleRandomizedSet() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Remove(2))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.Remove(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.GetRandom())
	// Output:
	// true
	// false
	// true
	// 2
	// true
	// false
	// 2
}
