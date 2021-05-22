package main

import "fmt"

func ExampleSolution() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	s := Constructor(head)
	cnt := [4]int{}
	for i := 0; i < 10000; i++ {
		cnt[s.GetRandom()]++
	}

	ok := true
	for i := 1; i <= 3; i++ {
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
