package main

import "sort"

/*
  You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

  Return the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

  You may assume that you have an infinite number of each kind of coin.

  Example 1:
    Input: coins = [1,2,5], amount = 11
    Output: 3
    Explanation: 11 = 5 + 5 + 1

  Example 2:
    Input: coins = [2], amount = 3
    Output: -1

  Example 3:
    Input: coins = [1], amount = 0
    Output: 0

  Example 4:
    Input: coins = [1], amount = 1
    Output: 1

  Example 5:
    Input: coins = [1], amount = 2
    Output: 2

  Constraints:
    1. 1 <= coins.length <= 12
    2. 1 <= coins[i] <= 2^31 - 1
    3. 0 <= amount <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/coin-change
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = 1<<31 - 1
		for _, coin := range coins {
			if coin > i {
				break
			}
			dp[i] = Min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == 1<<31-1 {
		return -1
	}
	return dp[amount]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
