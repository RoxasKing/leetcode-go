package leetcode

/*
  爱丽丝参与一个大致基于纸牌游戏 “21点” 规则的游戏，描述如下：

  爱丽丝以 0 分开始，并在她的得分少于 K 分时抽取数字。 抽取时，她从 [1, W] 的范围中随机获得一个整数作为分数进行累计，其中 W 是整数。 每次抽取都是独立的，其结果具有相同的概率。

  当爱丽丝获得不少于 K 分时，她就停止抽取数字。 爱丽丝的分数不超过 N 的概率是多少？

  提示：
    0 <= K <= N <= 10000
    1 <= W <= 10000
    如果答案与正确答案的误差不超过 10^-5，则该答案将被视为正确答案通过。
    此问题的判断限制时间已经减少。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/new-21-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func new21Game(N int, K int, W int) float64 {
	if K == 0 {
		return 1
	}
	dp := make([]float64, K+W)
	for i := K; i <= N && i < K+W; i++ {
		dp[i] = 1
	}
	dp[K-1] = float64(Min(N-K+1, W)) / float64(W)
	for i := K - 2; i >= 0; i-- {
		dp[i] = dp[i+1] + (dp[i+1+W]-dp[i+1])/float64(W)
	}
	return dp[0]
}

func new21Game2(N int, K int, W int) float64 {
	if K == 0 || N >= K+W {
		return 1
	}
	sum := make([]float64, K+W)
	sum[0] = 1
	for i := 1; i <= W; i++ {
		t := Min(i-1, K-1)
		sum[i] = sum[i-1] + sum[t]/float64(W)
	}
	for i := W + 1; i < len(sum); i++ {
		t := Min(i-1, K-1)
		sum[i] = sum[i-1] + (sum[t]-sum[i-W-1])/float64(W)
	}
	return (sum[N] - sum[K-1]) / (sum[K+W-1] - sum[K-1])
}
