package main

import (
	"fmt"
	"math"
)

func ExampleSolution() {
	so := Constructor([]int{1, 3})
	cnt := [2]int{}
	for i := 0; i < 10000; i++ {
		cnt[so.PickIndex()]++
	}
	if math.Abs(float64(cnt[0])-2500) > 10 {
		fmt.Println(cnt[0])
		fmt.Println("Test Failed!")
	} else {
		fmt.Println("Test Successed!")
	}
	// Output:
	// Test Successed!
}
