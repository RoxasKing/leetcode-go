package main

// Difficulty:
// Hard

// Tags:
// Math
// Enumeration

func consecutiveNumbersSum(n int) int {
	sum := func(i int) int { return (1 + i) * i / 2 }
	o := 1
	for i := 2; sum(i) <= n; i++ {
		if (n-sum(i))%i == 0 {
			o++
		}
	}
	return o
}
