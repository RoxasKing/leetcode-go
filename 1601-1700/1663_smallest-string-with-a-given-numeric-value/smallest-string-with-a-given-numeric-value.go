package main

// Difficulty:
// Medium

// Tags:
// Greedy

func getSmallestString(n int, k int) string {
	arr := make([]byte, n)
	for i := range arr {
		arr[i] = 'a'
	}
	for i, k := n-1, k-n; k > 0; i-- {
		x := Min(25, k)
		arr[i] += byte(x)
		k -= x
	}
	return string(arr)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
