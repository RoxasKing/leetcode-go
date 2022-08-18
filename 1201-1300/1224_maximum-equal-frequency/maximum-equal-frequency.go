package main

// Difficulty:
// Hard

// Tags:
// Stack
// Hash

func maxEqualFreq(nums []int) int {
	stk := []int{0}
	top := func() int { return len(stk) - 1 }
	freq := [1e5 + 1]int{}
	o := 0
	for i, x := range nums {
		if freq[x]++; top() < freq[x] {
			stk = append(stk, 0)
		}
		if stk[freq[x]-1] > 0 {
			stk[freq[x]-1]--
		}
		stk[freq[x]]++
		if i == top()*stk[top()] ||
			stk[top()] == 1 && i+1 == top()+(top()-1)*stk[top()-1] ||
			top() == 1 && i+1 == stk[top()] {
			o = i + 1
		}
	}
	return o
}
