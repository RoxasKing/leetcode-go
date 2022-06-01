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
	n := len(matchsticks)
	vis := make([]bool, 1<<n)
	mask := 1<<n - 1
	cur, cnt := 0, 0
	var backtrack func() bool
	backtrack = func() bool {
		if cnt == 4 {
			return true
		}
		if vis[mask] {
			return false
		}
		vis[mask] = true
		for i := 0; i < n; i++ {
			if (mask>>i)&1 == 0 || cur+matchsticks[i] > avg {
				continue
			}
			mask ^= 1 << i
			if cur+matchsticks[i] < avg {
				cur += matchsticks[i]
				if backtrack() {
					return true
				}
				cur -= matchsticks[i]
			} else {
				tmp := cur
				cur, cnt = 0, cnt+1
				if backtrack() {
					return true
				}
				cur, cnt = tmp, cnt-1
			}
			mask |= 1 << i
		}
		return false
	}
	return backtrack()
}
