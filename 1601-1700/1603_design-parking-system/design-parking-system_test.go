package main

import "fmt"

func ExampleParkingSystem() {
	ps := Constructor(1, 1, 0)
	fmt.Println(ps.AddCar(1))
	fmt.Println(ps.AddCar(2))
	fmt.Println(ps.AddCar(3))
	fmt.Println(ps.AddCar(1))
	// Output:
	// true
	// true
	// false
	// false
}
