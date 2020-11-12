package main

import (
	"fmt"
	"math"
)

func ExampleRandomizedCollection() {
	collection := Constructor()
	fmt.Println(collection.Insert(1))
	fmt.Println(collection.Insert(1))
	fmt.Println(collection.Insert(2))
	times := 100000
	count := 0
	for i := 0; i < times; i++ {
		if collection.GetRandom() == 1 {
			count++
		}
	}
	if math.Abs((float64(count)/float64(times))*3-2) > 0.01 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
	fmt.Println(collection.Remove(1))
	count = 0
	for i := 0; i < times; i++ {
		if collection.GetRandom() == 1 {
			count++
		}
	}
	if math.Abs((float64(count)/float64(times))*2-1) > 0.01 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
	fmt.Println(collection.Insert(3))
	fmt.Println(collection.Remove(2))
	count = 0
	for i := 0; i < times; i++ {
		if collection.GetRandom() == 1 {
			count++
		}
	}
	if math.Abs((float64(count)/float64(times))*2-1) > 0.01 {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
	// Output:
	// true
	// false
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}
