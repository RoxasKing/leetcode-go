package main

/*
  一条包含字母 A-Z 的消息通过以下方式进行了编码：
    'A' -> 1
    'B' -> 2
    ...
    'Z' -> 26
  给定一个只包含数字的非空字符串，请计算解码方法的总数。
  题目数据保证答案肯定是一个 32 位的整数。

  示例 1：
    输入：s = "12"
    输出：2
    解释：它可以解码为 "AB"（1 2）或者 "L"（12）。

  示例 2：
    输入：s = "226"
    输出：3
    解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。

  示例 3：
    输入：s = "0"
    输出：0

  示例 4：
    输入：s = "1"
    输出：1

  示例 5：
    输入：s = "2"
    输出：1

  提示：
    1 <= s.length <= 100
    s 只包含数字，并且可能包含前导零。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/decode-ways
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < n; i++ {
		if s[i] != '0' {
			dp[i+1] += dp[i]
		}
		if s[i-1] == '1' || s[i-1] == '2' && s[i] <= '6' {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[n]
}
