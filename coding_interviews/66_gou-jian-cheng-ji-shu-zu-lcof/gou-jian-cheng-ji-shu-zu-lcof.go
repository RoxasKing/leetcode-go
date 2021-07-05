package main

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	out := make([]int, n)
	out[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		out[i] = out[i+1] * a[i+1]
	}
	mul := 1
	for i := 0; i < n; i++ {
		out[i] *= mul
		mul *= a[i]
	}
	return out
}
