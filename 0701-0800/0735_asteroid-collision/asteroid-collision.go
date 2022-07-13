package main

// Difficulty:
// Medium

// Tags:
// Stack

func asteroidCollision(asteroids []int) []int {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	for _, x := range asteroids {
		for ; len(stk) > 0 && stk[top()] > 0 && x < 0; stk = stk[:top()] {
			if abs(stk[top()]) == abs(x) {
				x = 0
			} else if abs(stk[top()]) > abs(x) {
				x = stk[top()]
			}
		}
		if x != 0 {
			stk = append(stk, x)
		}
	}
	return stk
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
