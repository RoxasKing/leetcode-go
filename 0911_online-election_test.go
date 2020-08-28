package main

import "fmt"

func ExampleTopVotedCandidate() {
	t := NewTopVotedCandidate(
		[]int{0, 1, 1, 0, 0, 1, 0},
		[]int{0, 5, 10, 15, 20, 25, 30},
	)
	fmt.Println(t.Q(3))
	fmt.Println(t.Q(12))
	fmt.Println(t.Q(25))
	fmt.Println(t.Q(15))
	fmt.Println(t.Q(24))
	fmt.Println(t.Q(8))
	// Output:
	// 0
	// 1
	// 1
	// 0
	// 0
	// 1
}
