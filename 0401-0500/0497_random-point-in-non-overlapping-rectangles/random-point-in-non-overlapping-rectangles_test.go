package main

import "fmt"

func ExampleSolution() {
	s := Constructor([][]int{{-2, -2, -1, -1}, {1, 0, 3, 0}})
	cnt := make(map[int]int)
	for i := 0; i < 70000; i++ {
		res := s.Pick()
		i, j := res[0], res[1]
		cnt[i*100+j]++
	}
	ok := true
	for _, c := range cnt {
		if c-10000 > 1000 || 10000-c > 1000 {
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
