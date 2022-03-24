package main

// Difficulty:
// Medium

// Tags:
// Math
// Greedy

func brokenCalc(startValue int, target int) int {
	cnt := 0
	for startValue < target {
		cnt++
		if target&1 == 1 {
			target++
		} else {
			target >>= 1
		}
	}
	return cnt + startValue - target
}
