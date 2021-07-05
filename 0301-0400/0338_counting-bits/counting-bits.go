package main

// Tags:
// Dynamic Programming
func countBits(num int) []int {
	out := make([]int, num+1)
	for i := 1; i <= num; i++ {
		out[i] = out[i>>1]
		if i&1 == 1 {
			out[i]++
		}
	}
	return out
}
