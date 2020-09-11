package main

/*
  给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。

  为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -228 到 228 - 1 之间，最终结果不会超过 231 - 1 。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/4sum-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash
func fourSumCount(A []int, B []int, C []int, D []int) int {
	var out int
	sum := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			sum[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			out += sum[0-c-d]
		}
	}
	return out
}
