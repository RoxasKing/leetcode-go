package main

import "fmt"

func ExampleMovingAverage() {
	ma := Constructor(3)
	fmt.Printf("%.5f\n", ma.Next(1))
	fmt.Printf("%.5f\n", ma.Next(10))
	fmt.Printf("%.5f\n", ma.Next(3))
	fmt.Printf("%.5f\n", ma.Next(5))
	// Output:
	// 1.00000
	// 5.50000
	// 4.66667
	// 6.00000
}
