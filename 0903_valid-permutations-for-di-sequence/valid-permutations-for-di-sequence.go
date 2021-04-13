package main

/*
  We are given S, a length n string of characters from the set {'D', 'I'}. (These letters stand for "decreasing" and "increasing".)

  A valid permutation is a permutation P[0], P[1], ..., P[n] of integers {0, 1, ..., n}, such that for all i:
    1. If S[i] == 'D', then P[i] > P[i+1], and;
    2. If S[i] == 'I', then P[i] < P[i+1].
  How many valid permutations are there?  Since the answer may be large, return your answer modulo 10^9 + 7.

  Example 1:
    Input: "DID"
    Output: 5
    Explanation:
      The 5 valid permutations of (0, 1, 2, 3) are:
      (1, 0, 3, 2)
      (2, 0, 3, 1)
      (2, 1, 3, 0)
      (3, 0, 2, 1)
      (3, 1, 2, 0)

  Note:
    1. 1 <= S.length <= 200
    2. S consists only of characters from the set {'D', 'I'}.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-permutations-for-di-sequence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func numPermsDISequence(S string) int {
	mod := int(1e9 + 7)
	n := len(S)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			l, r := 0, j // S[i-1] == 'I'
			if S[i-1] == 'D' {
				l, r = j, i
			}
			for k := l; k < r; k++ {
				dp[i][j] += dp[i-1][k]
				dp[i][j] %= mod
			}
		}
	}

	out := 0
	for _, cnt := range dp[n] {
		out += cnt
		out %= mod
	}
	return out
}

// Dynamic Programming (O(S.length) additional memory space)
func numPermsDISequence2(S string) int {
	mod := int(1e9 + 7)
	n := len(S)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			l, r := 0, j // S[i-1] == 'I'
			if S[i-1] == 'D' {
				l, r = j, i
			}
			dp[1][j] = 0
			for k := l; k < r; k++ {
				dp[1][j] += dp[0][k]
				dp[1][j] %= mod
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}

	out := 0
	for j := range dp[0] {
		out += dp[0][j]
		out %= mod
	}
	return out
}
