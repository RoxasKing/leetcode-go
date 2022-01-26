package main

import "fmt"

func ExampleStockPrice() {
	sp := Constructor()
	sp.Update(1, 10)
	sp.Update(2, 5)
	fmt.Println(sp.Current())
	fmt.Println(sp.Maximum())
	sp.Update(1, 3)
	fmt.Println(sp.Maximum())
	sp.Update(4, 2)
	fmt.Println(sp.Minimum())
	// Output:
	// 5
	// 10
	// 5
	// 2
}
