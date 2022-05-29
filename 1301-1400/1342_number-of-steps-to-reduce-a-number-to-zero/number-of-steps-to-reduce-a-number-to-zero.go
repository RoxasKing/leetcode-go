package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

func numberOfSteps(num int) int {
	o := 0
	for ; num > 0; o++ {
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
	}
	return o
}
