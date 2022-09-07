package main

// Difficulty:
// Easy

// Tags:
// Monotonic Stack

func finalPrices(prices []int) []int {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	n := len(prices)
	o := make([]int, n)
	for j, p := range prices {
		for ; len(stk) > 0 && prices[stk[top()]] >= p; stk = stk[:top()] {
			i := stk[top()]
			o[i] = prices[i] - p
		}
		o[j] = p
		stk = append(stk, j)
	}
	return o
}
