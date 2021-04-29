package main

/*
  Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.

  If there is no non-empty subarray with sum at least K, return -1.

  Example 1:
    Input: A = [1], K = 1
    Output: 1

  Example 2:
    Input: A = [1,2], K = 4
    Output: -1

  Example 3:
    Input: A = [2,-1,2], K = 3
    Output: 3

  Note:
    1. 1 <= A.length <= 50000
    2. -10 ^ 5 <= A[i] <= 10 ^ 5
    3. 1 <= K <= 10 ^ 9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shortest-subarray-with-sum-at-least-k
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum + Sliding Window
func shortestSubarray(A []int, K int) int {
	n := len(A)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + A[i]
	}
	out := 1<<31 - 1

	for i, q := 0, []int{}; i <= n; i++ {
		for len(q) > 0 && pSum[i] <= pSum[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		for ; len(q) > 0 && pSum[i]-pSum[q[0]] >= K; q = q[1:] {
			out = Min(out, i-q[0])
		}
		q = append(q, i)
	}

	if out == 1<<31-1 {
		return -1
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
