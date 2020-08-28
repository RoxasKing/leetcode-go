package main

/*
  一条包含字母 A-Z 的消息通过以下方式进行了编码：
    'A' -> 1
    'B' -> 2
    ...
    'Z' -> 26
  给定一个只包含数字的非空字符串，请计算解码方法的总数。
*/

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-2] == '1' || s[i-2] == '2' && s[i-1] <= '6' {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
