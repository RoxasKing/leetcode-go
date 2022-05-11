package main

// Difficulty:
// Medium

// Tags:
// Monotonic Stack

func find132pattern(nums []int) bool {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	for num2, i := -1<<31, len(nums)-1; i >= 0; i-- {
		if nums[i] < num2 {
			return true
		}
		for ; len(stk) > 0 && stk[top()] < nums[i]; stk = stk[:top()] {
			num2 = stk[top()]
		}
		if nums[i] > num2 {
			stk = append(stk, nums[i])
		}
	}
	return false
}
