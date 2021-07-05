package main

func validMountainArray(A []int) bool {
	n := len(A)
	if n < 3 {
		return false
	}
	l, r := 0, n-1
	for l+1 < n-1 && A[l+1] > A[l] {
		l++
	}
	for 0 < r-1 && A[r-1] > A[r] {
		r--
	}
	return l == r
}
