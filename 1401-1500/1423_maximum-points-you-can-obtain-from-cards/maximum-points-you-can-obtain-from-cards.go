package main

// Tags:
// Sliding Window
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	sum := 0
	for i := n - k; i < n; i++ {
		sum += cardPoints[i]
	}
	out := sum
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
		sum -= cardPoints[i-k+n]
		out = Max(out, sum)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
