package main

// Difficulty:
// Medium

// Tags:
// Stack

func wiggleMaxLength(nums []int) int {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	for _, x := range nums {
		if len(stk) > 0 && stk[top()] == x {
			continue
		}
		if len(stk) < 2 ||
			stk[top()-1] < stk[top()] && stk[top()] > x ||
			stk[top()-1] > stk[top()] && stk[top()] < x {
			stk = append(stk, x)
			continue
		}
		stk[top()] = x
	}
	return len(stk)
}
