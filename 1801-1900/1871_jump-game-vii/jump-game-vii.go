package main

// Tags:
// Greedy Algoritym

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] != '0' {
		return false
	}

	l, r := minJump, Min(maxJump, n-1)
	for l <= r && r < n-1 {
		nextL, nextR := 1<<31-1, -1<<31
		for i := l; i <= Min(r, n-1); i++ {
			if s[i] == '0' {
				nextL, nextR = Min(nextL, i+minJump), i+maxJump
			}
		}
		l, r = Max(r+1, nextL), nextR
	}
	return l <= n-1 && n-1 <= r
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
