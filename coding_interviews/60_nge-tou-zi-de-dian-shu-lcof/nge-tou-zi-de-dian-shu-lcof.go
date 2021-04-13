package main

/*
  把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。

  你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

  示例 1:
    输入: 1
    输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]

  示例 2:
    输入: 2
    输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0.05556,0.02778]

  限制：
    1 <= n <= 11

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func dicesProbability(n int) []float64 {
	cnt, tmp := make([]int, 6*n+1), make([]int, 6*n+1)
	for i := 1; i <= 6; i++ {
		cnt[i] = 1
	}
	t := 6
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			tmp[j] = 0
		}
		for j := i; j <= 6*i; j++ {
			tmp[j] = 0
			for k := 1; k <= 6 && k < j; k++ {
				tmp[j] += cnt[j-k]
			}
		}
		cnt, tmp = tmp, cnt
		t *= 6
	}
	cnt = cnt[n:]
	out := make([]float64, len(cnt))
	for i := range out {
		out[i] = float64(cnt[i]) / float64(t)
	}
	return out
}
