package main

/*
  Alice plays the following game, loosely based on the card game "21".

  Alice starts with 0 points, and draws numbers while she has less than K points.  During each draw, she gains an integer number of points randomly from the range [1, W], where W is an integer.  Each draw is independent and the outcomes have equal probabilities.

  Alice stops drawing numbers when she gets K or more points.  What is the probability that she has N or less points?

  Example 1:
    Input: N = 10, K = 1, W = 10
    Output: 1.00000
    Explanation:  Alice gets a single card, then stops.

  Example 2:
    Input: N = 6, K = 1, W = 10
    Output: 0.60000
    Explanation:  Alice gets a single card, then stops.
    In 6 out of W = 10 possibilities, she is at or below N = 6 points.

  Example 3:
    Input: N = 21, K = 17, W = 10
    Output: 0.73278

  Note:
    1. 0 <= K <= N <= 10000
    2. 1 <= W <= 10000
    3. Answers will be accepted as correct if they are within 10^-5 of the correct answer.
    4. The judging time limit has been reduced for this question.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/new-21-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math + Dynamic Programming
func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1
	}
	dp := make([]float64, K+W)
	for i := K; i <= N && i <= K-1+W; i++ {
		dp[i] = 1
	}
	dp[K-1] = float64(Min(N, K-1+W)-(K-1)) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] - (dp[i+W+1]-dp[i+1])/float64(W)
	}
	return dp[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
