package main

// Tags:
// Sliding Window
func equalSubstring(s string, t string, maxCost int) int {
	out := 0
	for l, r, cost := 0, 0, 0; r < len(s); r++ {
		cost += Abs(int(s[r]) - int(t[r]))
		for l <= r && cost > maxCost {
			cost -= Abs(int(s[l]) - int(t[l]))
			l++
		}
		out = Max(out, r+1-l)
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
