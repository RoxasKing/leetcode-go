package main

// Difficulty:
// Medium

// Tags:
// Bitmask
// Bit Manipulation
// Backtracking

func makesquare(matchsticks []int) bool {
	sum := 0
	for _, x := range matchsticks {
		sum += x
	}
	if sum%4 != 0 {
		return false
	}
	avg := sum / 4
	n := 1 << len(matchsticks)
	vis := make([]bool, n)
	mask := n - 1
	var cur, cnt int
	var backtrack func() bool
	backtrack = func() bool {
		if cnt == 4 {
			return true
		}
		if vis[mask] {
			return false
		}
		vis[mask] = true
		for i, x := range matchsticks {
			if (mask>>i)&1 == 0 || cur+x > avg {
				continue
			}
			mask ^= 1 << i
			if cur += x; cur == avg {
				cur, cnt = 0, cnt+1
			}
			if backtrack() {
				return true
			}
			if cur -= x; cur < 0 {
				cur, cnt = avg-x, cnt-1
			}
			mask ^= 1 << i
		}
		return false
	}
	return backtrack()
}
