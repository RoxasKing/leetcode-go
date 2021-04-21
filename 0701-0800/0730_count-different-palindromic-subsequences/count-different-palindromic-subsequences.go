package main

/*
  给定一个字符串 S，找出 S 中不同的非空回文子序列个数，并返回该数字与 10^9 + 7 的模。
  通过从 S 中删除 0 个或多个字符来获得子序列。
  如果一个字符序列与它反转后的字符序列一致，那么它是回文字符序列。
  如果对于某个  i，A_i != B_i，那么 A_1, A_2, ... 和 B_1, B_2, ... 这两个字符序列是不同的。

  示例 1：
    输入：
    S = 'bccb'
    输出：6
    解释：
    6 个不同的非空回文子字符序列分别为：'b', 'c', 'bb', 'cc', 'bcb', 'bccb'。
    注意：'bcb' 虽然出现两次但仅计数一次。

  示例 2：
    输入：
    S = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
    输出：104860361
    解释：
    共有 3104860382 个不同的非空回文子序列，对 10^9 + 7 取模为 104860361。

  提示：
    字符串 S 的长度将在[1, 1000]范围内。
    每个字符 S[i] 将会是集合 {'a', 'b', 'c', 'd'} 中的某一个。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-different-palindromic-subsequences
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func countPalindromicSubsequences(S string) int {
	const mod = 1e9 + 7
	n := len(S)
	dp := make([][][]int, 4)
	for k := range dp {
		dp[k] = make([][]int, n)
		for i := range dp[k] {
			dp[k][i] = make([]int, n)
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			for k := 0; k < 4; k++ {
				c := byte(k) + 'a'
				if i == j && S[i] == c {
					dp[k][i][j] = 1
				} else if i != j {
					if S[i] != c {
						dp[k][i][j] = dp[k][i+1][j]
					} else if S[j] != c {
						dp[k][i][j] = dp[k][i][j-1]
					} else {
						dp[k][i][j] = 2
						if i+1 == j {
							continue
						}
						for m := 0; m < 4; m++ {
							dp[k][i][j] += dp[m][i+1][j-1]
							dp[k][i][j] %= mod
						}
					}
				}
			}
		}
	}
	var count int
	for k := 0; k < 4; k++ {
		count += dp[k][0][n-1]
		count %= mod
	}
	return count
}

// Dynamic Programming
func countPalindromicSubsequences2(S string) int {
	n := len(S)
	prev := makeArrrays(n, 4, -1)
	next := makeArrrays(n, 4, -1)
	memo := makeArrrays(n, n, 0)

	charArr := []byte(S)

	last := makeArray(4, -1)
	for i := 0; i < n; i++ {
		last[charArr[i]-'a'] = i
		for k := 0; k < 4; k++ {
			prev[i][k] = last[k]
		}
	}

	fillArray(last, -1)
	for i := n - 1; i >= 0; i-- {
		last[charArr[i]-'a'] = i
		for k := 0; k < 4; k++ {
			next[i][k] = last[k]
		}
	}

	return dp(memo, prev, next, 0, n-1) - 1
}

func dp(memo, prev, next [][]int, i, j int) int {
	if memo[i][j] > 0 {
		return memo[i][j]
	}
	out := 1
	if i <= j {
		for k := 0; k < 4; k++ {
			i0, j0 := next[i][k], prev[j][k]
			if i <= i0 && i0 <= j {
				out++
			}
			if -1 < i0 && i0 < j0 {
				out += dp(memo, prev, next, i0+1, j0-1)
			}
			if out >= 1e9+7 {
				out -= 1e9 + 7
			}
		}
	}
	memo[i][j] = out
	return out
}

func makeArray(n, initVal int) []int {
	arr := make([]int, n)
	fillArray(arr, initVal)
	return arr
}

func makeArrrays(row, col, initVal int) [][]int {
	arr := make([][]int, row)
	for i := range arr {
		arr[i] = make([]int, col)
		if initVal != 0 {
			fillArray(arr[i], initVal)
		}
	}
	return arr
}

func fillArray(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}
