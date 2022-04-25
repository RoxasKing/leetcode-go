package main

import "fmt"

func ExampleUndergroundSystem() {
	undergroundSystem := Constructor()
	undergroundSystem.CheckIn(45, "Leyton", 3)
	undergroundSystem.CheckIn(32, "Paradise", 8)
	undergroundSystem.CheckIn(27, "Leyton", 10)
	undergroundSystem.CheckOut(45, "Waterloo", 15)
	undergroundSystem.CheckOut(27, "Waterloo", 20)
	undergroundSystem.CheckOut(32, "Cambridge", 22)
	fmt.Println(undergroundSystem.GetAverageTime("Paradise", "Cambridge"))
	fmt.Println(undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckIn(10, "Leyton", 24)
	fmt.Println(undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	undergroundSystem.CheckOut(10, "Waterloo", 38)
	fmt.Println(undergroundSystem.GetAverageTime("Leyton", "Waterloo"))
	// Output:
	// 14
	// 11
	// 11
	// 12
}
