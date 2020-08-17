package main

/*
  给出一些不同颜色的盒子，盒子的颜色由数字表示，即不同的数字表示不同的颜色。
  你将经过若干轮操作去去掉盒子，直到所有的盒子都去掉为止。每一轮你可以移除具有相同颜色的连续 k 个盒子（k >= 1），这样一轮之后你将得到 k*k 个积分。
  当你将所有盒子都去掉之后，求你能获得的最大积分和。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-boxes
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeBoxes(boxes []int) int {
	dp := [100][100][100]int{}
	var cal func(int, int, int) int
	cal = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if dp[l][r][k] != 0 {
			return dp[l][r][k]
		}
		for l < r && boxes[r-1] == boxes[r] {
			r--
			k++
		}
		dp[l][r][k] = cal(l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				dp[l][r][k] = Max(dp[l][r][k], cal(l, i, k+1)+cal(i+1, r-1, 0))
			}
		}
		return dp[l][r][k]
	}
	return cal(0, len(boxes)-1, 0)
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
