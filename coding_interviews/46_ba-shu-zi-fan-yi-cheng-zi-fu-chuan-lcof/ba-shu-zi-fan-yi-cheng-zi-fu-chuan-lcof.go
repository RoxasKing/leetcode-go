package main

import "strconv"

/*
  给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

  示例 1:
    输入: 12258
    输出: 5
    解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"

  提示：
    0 <= num < 2^31

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Recursion
func translateNum(num int) int {
	str := strconv.Itoa(num)
	out := 0
	dfs(str, 0, &out)
	return out
}

func dfs(str string, i int, out *int) {
	if i == len(str) {
		*out++
		return
	}

	dfs(str, i+1, out)

	if i+1 < len(str) && (str[i] == '1' || str[i] == '2' && str[i+1] <= '5') {
		dfs(str, i+2, out)
	}
}

// Dynamic Programming
func translateNum2(num int) int {
	arr := []int{}
	for ; num > 0; num /= 10 {
		arr = append(arr, num%10)
	}
	reverse(arr)

	n := len(arr)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		dp[i+1] = dp[i]
		if i-1 >= 0 && (arr[i-1] == 1 || arr[i-1] == 2 && arr[i] <= 5) {
			dp[i+1] += dp[i-1]
		}
	}

	return dp[n]
}

func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n>>1; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}
