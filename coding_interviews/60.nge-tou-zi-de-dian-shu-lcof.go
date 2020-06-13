package codinginterviews

/*
  把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

  你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

  限制：
    1 <= n <= 11

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSumII(n int) []float64 {
	times, newTimes := make([]int, 6*n+1), make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		times[i] = 1
	}
	sumTimes := 6
	for k := 2; k <= n; k++ {
		for i := 0; i < k; i++ {
			newTimes[i] = 0
		}
		for i := k; i <= 6*k; i++ {
			newTimes[i] = 0
			for j := 1; j <= i && j <= 6; j++ {
				newTimes[i] += times[i-j]
			}
		}
		times, newTimes = newTimes, times
		sumTimes *= 6
	}
	times = times[n:]
	out := make([]float64, 6*n-n+1)
	for i := range out {
		out[i] = float64(times[i]) / float64(sumTimes)
	}
	return out
}
