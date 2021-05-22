package main

import "fmt"

func ExampleSolution() {
	s := Constructor([]int{1, 2, 3, 3, 3})
	cnt := [5]int{}
	for i := 0; i < 10000; i++ {
		cnt[s.Pick(3)]++
	}
	ok := true
	for i := 2; i <= 4; i++ {
		if cnt[i]-3333 > 100 || cnt[i]-3333 < -100 {
			ok = false
			break
		}
	}
	if ok {
		fmt.Println("Test succeed!")
	} else {
		fmt.Println("Test failed!")
	}
	// Output:
	// Test succeed!
}
