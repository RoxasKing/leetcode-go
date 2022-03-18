package main

// Difficulty:
// Medium

// Tags:
// Stack

func validateStackSequences(pushed []int, popped []int) bool {
	k, n := 0, len(pushed)
	for i, stk := 0, []int{}; i < n; i++ {
		stk = append(stk, pushed[i])
		for ; len(stk) > 0 && stk[len(stk)-1] == popped[k]; k++ {
			stk = stk[:len(stk)-1]
		}
	}
	return k == n
}
