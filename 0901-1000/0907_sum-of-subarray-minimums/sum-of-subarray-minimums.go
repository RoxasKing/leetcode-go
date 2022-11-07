package main

// Difficulty:
// Medium

// Tags:
// Monotonic Stack

func sumSubarrayMins(arr []int) int {
	const mod int = 1e9 + 7
	stk := []int{-1}
	top := func() int { return len(stk) - 1 }
	o := 0
	for i, x := range arr {
		for len(stk) > 1 && arr[stk[top()]] >= x {
			k := stk[top()]
			stk = stk[:top()]
			v := arr[k]
			o = (o + (i-k)*(k-stk[top()])*v) % mod
		}
		stk = append(stk, i)
	}
	for i := 1; i < len(stk); i++ {
		o = (o + (len(arr)-stk[i])*(stk[i]-stk[i-1])*arr[stk[i]]) % mod
	}
	return o
}
