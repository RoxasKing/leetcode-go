package main

// Tags:
// Dynamic Programming
func longestMountain(a []int) (ans int) {
	n := len(a)
	L := make([]int, n)
	for i := 1; i < n-1; i++ {
		if a[i] > a[i-1] {
			L[i] = L[i-1] + 1
		}
	}
	R := make([]int, n)
	for i := n - 2; i > 0; i-- {
		if a[i] > a[i+1] {
			R[i] = R[i+1] + 1
		}
	}
	for i := 1; i < n-1; i++ {
		if L[i] > 0 && R[i] > 0 {
			ans = Max(ans, L[i]+R[i]+1)
		}
	}
	return
}

func longestMountain2(a []int) (ans int) {
	n := len(a)
	for l, r := 0, 1; l < n-2; l, r = r, r+1 {
		if a[l] >= a[r] {
			continue
		}
		for r < n-1 && a[r] < a[r+1] {
			r++
		}
		if r == n-1 || a[r] == a[r+1] {
			r++
			continue
		}
		for r < n-1 && a[r] > a[r+1] {
			r++
		}
		ans = Max(ans, r+1-l)
	}
	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
