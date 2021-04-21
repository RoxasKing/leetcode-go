package main

/*
  Alice and Bob take turns playing a game, with Alice starting first.

  Initially, there is a number n on the chalkboard. On each player's turn, that player makes a move consisting of:

    1. Choosing any x with 0 < x < n and n % x == 0.
    2. Replacing the number n on the chalkboard with n - x.

  Also, if a player cannot make a move, they lose the game.

  Return true if and only if Alice wins the game, assuming both players play optimally.

  Example 1:
    Input: n = 2
    Output: true
    Explanation: Alice chooses 1, and Bob has no more moves.

  Example 2:
    Input: n = 3
    Output: false
    Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.

  Constraints:
    1 <= n <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/divisor-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func divisorGame(n int) bool {
	dp := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 && !dp[i-j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// Math
func divisorGame2(N int) bool {
	return N&1 == 0
}
