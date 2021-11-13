package main

// Difficulty:
// Medium

// Tags:
// Monotonic Stack

func dailyTemperatures(temperatures []int) []int {
	out := make([]int, len(temperatures))
	stk := []int{}
	for j := range temperatures {
		for len(stk) > 0 && temperatures[j] > temperatures[stk[len(stk)-1]] {
			i := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			out[i] = j - i
		}
		stk = append(stk, j)
	}
	return out
}
