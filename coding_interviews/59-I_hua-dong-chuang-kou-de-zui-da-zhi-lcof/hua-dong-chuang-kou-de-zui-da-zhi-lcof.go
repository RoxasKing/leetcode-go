package main

// Tags:
// Monotonous Queue
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	out := make([]int, 0, n+1-k)
	stk := []int{}
	for i, num := range nums {
		for len(stk) > 0 && stk[len(stk)-1] < num {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, num)
		if i > k-1 && stk[0] == nums[i-k] {
			stk = stk[1:]
		}
		if i >= k-1 {
			out = append(out, stk[0])
		}
	}
	return out
}
