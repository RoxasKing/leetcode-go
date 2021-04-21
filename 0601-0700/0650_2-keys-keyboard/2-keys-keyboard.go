package main

/*
  Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:
    1. Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
    2. Paste: You can paste the characters which are copied last time.
  Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.

  Example 1:
    Input: 3
    Output: 3
    Explanation:
    Intitally, we have one character 'A'.
    In step 1, we use Copy All operation.
    In step 2, we use Paste operation to get 'AA'.
    In step 3, we use Paste operation to get 'AAA'.

  Note:
    The n will be in the range [1, 1000].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/2-keys-keyboard
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS
func minSteps(n int) int {
	out := n
	dfs(n, 1, 0, 0, &out)
	return out
}

func dfs(n, cur, steps, prevStep int, out *int) {
	if steps >= *out {
		return
	}

	if cur == n {
		*out = Min(*out, steps)
		return
	}

	if cur+cur <= n {
		dfs(n, cur+cur, steps+2, cur, out)
	}

	if prevStep > 0 && cur+prevStep <= n {
		dfs(n, cur+prevStep, steps+1, prevStep, out)
	}
}

// Dynamic Programming
func minSteps2(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = 1<<31 - 1
		for j := 1; j < i; j++ {
			if i%j == 0 {
				dp[i] = Min(dp[i], dp[j]+i/j)
			}
		}
	}
	return dp[n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
