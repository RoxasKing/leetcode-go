package Algorithm

func CutInRecursive(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -2147483648
	for i := 1; i <= n; i++ {
		q = max(q, p[i-1]+CutInRecursive(p, n-i))
	}
	return q
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
