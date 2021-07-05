package main

// Tags:
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
