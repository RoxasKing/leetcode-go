package main

import "fmt"

func ExampleSolution() {
	c := Constructor(4, []int{2})
	cnt := [4]int{}
	for i := 0; i < 10000; i++ {
		cnt[c.Pick()]++
	}
	ok := true
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		if cnt[i]-3333 > 100 || 3333-cnt[i] > 100 {
			ok = false
			break
		}
	}
	if ok {
		fmt.Println("Test Succeed!")
	} else {
		fmt.Println("Test Failed!")
	}
	// Output:
	// Test Succeed!
}
