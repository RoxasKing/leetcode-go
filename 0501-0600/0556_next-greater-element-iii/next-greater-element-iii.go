package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Math
// Two Pointers

func nextGreaterElement(n int) int {
	a := []int{}
	for ; n > 0; n /= 10 {
		a = append(a, n%10)
	}
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	i := len(a) - 1
	for ; i > 0 && a[i-1] >= a[i]; i-- {
	}
	for l, r := i, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	if i == 0 {
		return -1
	}
	j := sort.Search(len(a)-i, func(j int) bool { return a[i+j] > a[i-1] }) + i
	a[i-1], a[j] = a[j], a[i-1]
	o := 0
	for _, x := range a {
		if o > (1<<31-1-x)/10 {
			return -1
		}
		o = o*10 + x
	}
	return o
}
