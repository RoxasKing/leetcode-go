package main

// Tags:
// Binary Search
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	a, b := 0, 1
	l, r := float64(0), float64(1)
	for r-l > 1e-9 {
		target := (l + r) / 2.0
		cnt, x, y := countSmaller(arr, n, target)
		if cnt < k {
			l = target
		} else {
			r = target
			a, b = x, y
		}
	}
	return []int{a, b}
}

func countSmaller(arr []int, n int, target float64) (int, int, int) {
	cnt, a, b := 0, 0, 1
	for i, j := -1, 1; j < n; j++ {
		for float64(arr[i+1])/float64(arr[j]) < target {
			i++
		}
		cnt += i + 1
		if i >= 0 && float64(a)/float64(b) < float64(arr[i])/float64(arr[j]) {
			a, b = arr[i], arr[j]
		}
	}
	return cnt, a, b
}
