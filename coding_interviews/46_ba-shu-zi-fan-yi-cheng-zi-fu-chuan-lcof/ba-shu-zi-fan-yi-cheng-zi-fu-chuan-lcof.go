package main

import "strconv"

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
