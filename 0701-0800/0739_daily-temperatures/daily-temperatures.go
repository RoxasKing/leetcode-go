package main

// Tags:
// Monotonic Stack
func dailyTemperatures(T []int) []int {
	n := len(T)
	out := make([]int, n)
	stk := []int{}
	for i := range T {
		for len(stk) > 0 && T[i] > T[stk[len(stk)-1]] {
			j := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			out[j] = i - j
		}
		stk = append(stk, i)
	}
	return out
}
