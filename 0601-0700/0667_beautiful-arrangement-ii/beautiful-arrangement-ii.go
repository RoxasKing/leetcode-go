package main

// Difficulty:
// Medium

// Tags:
// Math

func constructArray(n int, k int) []int {
	o := make([]int, 0, n)
	for i := 1; i < n-k; i++ {
		o = append(o, i)
	}
	l, r := n-k, n
	for ; l < r; l, r = l+1, r-1 {
		o = append(o, l, r)
	}
	if l == r {
		o = append(o, l)
	}
	return o
}
