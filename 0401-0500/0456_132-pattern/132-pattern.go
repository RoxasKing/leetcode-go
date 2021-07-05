package main

// Tags:
// Monotone Stack
func find132pattern(nums []int) bool {
	n := len(nums)
	stk := []int{}
	num2 := -1 << 31
	for i := n - 1; i >= 0; i-- {
		if nums[i] < num2 {
			return true
		}
		for len(stk) > 0 && stk[len(stk)-1] < nums[i] {
			num2 = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		if nums[i] > num2 {
			stk = append(stk, nums[i])
		}
	}
	return false
}
