package main

/*
  There is a tournament where n players are participating. The players are standing in a single row and are numbered from 1 to n based on their initial standing position (player 1 is the first player in the row, player 2 is the second player in the row, etc.).

  The tournament consists of multiple rounds (starting from round number 1). In each round, the ith player from the front of the row competes against the ith player from the end of the row, and the winner advances to the next round. When the number of players is odd for the current round, the player in the middle automatically advances to the next round.

  For example, if the row consists of players 1, 2, 4, 6, 7
    1. Player 1 competes against player 7.
    2. Player 2 competes against player 6.
    3. Player 4 automatically advances to the next round.

  After each round is over, the winners are lined back up in the row based on the original ordering assigned to them initially (ascending order).

  The players numbered firstPlayer and secondPlayer are the best in the tournament. They can win against any other player before they compete against each other. If any two other players compete against each other, either of them might win, and thus you may choose the outcome of this round.

  Given the integers n, firstPlayer, and secondPlayer, return an integer array containing two values, the earliest possible round number and the latest possible round number in which these two players will compete against each other, respectively.

  Example 1:
    Input: n = 11, firstPlayer = 2, secondPlayer = 4
    Output: [3,4]
    Explanation:
      One possible scenario which leads to the earliest round number:
      First round: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
      Second round: 2, 3, 4, 5, 6, 11
      Third round: 2, 3, 4
      One possible scenario which leads to the latest round number:
      First round: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
      Second round: 1, 2, 3, 4, 5, 6
      Third round: 1, 2, 4
      Fourth round: 2, 4

  Example 2:
    Input: n = 5, firstPlayer = 1, secondPlayer = 5
    Output: [1,1]
    Explanation: The players numbered 1 and 5 compete in the first round.
      There is no way to make them compete in any other round.

  Constraints:
    1. 2 <= n <= 28
    2. 1 <= firstPlayer < secondPlayer <= n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/the-earliest-and-latest-rounds-where-players-compete
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Porgramming

func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	if firstPlayer > secondPlayer {
		firstPlayer, secondPlayer = secondPlayer, firstPlayer
	}

	return dp(n, firstPlayer, secondPlayer)
}

func dp(n, f, s int) []int {
	if F[n][f][s] > 0 {
		return []int{F[n][f][s], G[n][f][s]}
	}

	if f+s == n+1 {
		return []int{1, 1}
	}

	if f+s > n+1 {
		out := dp(n, n+1-s, n+1-f)
		F[n][f][s], G[n][f][s] = out[0], out[1]
		return out
	}

	min, max := 1<<31-1, -1<<31
	halfN := (n + 1) >> 1

	if s <= halfN {
		for i := 0; i < f; i++ {
			for j := 0; j < s-f; j++ {
				out := dp(halfN, i+1, i+j+2)
				min = Min(min, out[0])
				max = Max(max, out[1])
			}
		}
	} else {
		sPrime := n + 1 - s
		mid := (n - 2*sPrime + 1) >> 1
		for i := 0; i < f; i++ {
			for j := 0; j < sPrime-f; j++ {
				out := dp(halfN, i+1, i+j+mid+2)
				min = Min(min, out[0])
				max = Max(max, out[1])
			}
		}
	}

	F[n][f][s] = min + 1
	G[n][f][s] = max + 1
	return []int{F[n][f][s], G[n][f][s]}
}

var F = [30][30][30]int{}
var G = [30][30][30]int{}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
