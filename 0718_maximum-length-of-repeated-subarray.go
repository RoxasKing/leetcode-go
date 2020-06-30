package leetcode

/*
  给两个整数数组 A 和 B ，返回两个数组中公共的、长度最长的子数组的长度。

  说明:
    1 <= len(A), len(B) <= 1000
    0 <= A[i], B[i] < 100
*/

// Sliding Window
func findLength(A []int, B []int) int {
	maxLen := func(indexA, indexB, limit int) int {
		var res, k int
		for i := 0; i < limit; i++ {
			if A[indexA+i] == B[indexB+i] {
				k++
			} else {
				k = 0
			}
			res = Max(res, k)
		}
		return res
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
