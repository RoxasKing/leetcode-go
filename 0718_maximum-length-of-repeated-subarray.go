package leetcode

/*
  给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

  说明:
    1 <= len(A), len(B) <= 1000
    0 <= A[i], B[i] < 100
*/

// Sliding Window
func findLength(A []int, B []int) int {
	maxLen := func(pA, pB, limit int) int {
		var max, count int
		for i := 0; i < limit; i++ {
			if A[pA+i] != B[pB+i] {
				count = 0
				continue
			}
			count++
			max = Max(max, count)
		}
		return max
	}
	var out int
	for i := range A {
		out = Max(out, maxLen(i, 0, Min(len(A), len(B)-i)))
	}
	for i := range B {
		out = Max(out, maxLen(0, i, Min(len(A)-i, len(B))))
	}
	return out
}

// Dynamic Programming
func findLength2(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	var out int
	for i := range A {
		for j := range B {
			if A[i] == B[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = 0
			}
			out = Max(out, dp[i+1][j+1])
		}
	}
	return out
}
