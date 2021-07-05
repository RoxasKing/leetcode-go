package main

// Tags:
// Math
func reverse(x int) int {
	out := 0
	for x != 0 {
		if out < (-1<<31-x%10)/10 || (1<<31-1-x%10)/10 < out {
			return 0
		}
		out = out*10 + x%10
		x /= 10
	}
	return out
}
