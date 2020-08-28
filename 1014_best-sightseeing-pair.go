package main

/*
  给定正整数数组 A，A[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的距离为 j - i。
  一对景点（i < j）组成的观光组合的得分为（A[i] + A[j] + i - j）：景点的评分之和减去它们两者之间的距离。
  返回一对观光景点能取得的最高分。

  提示：
    2 <= A.length <= 50000
    1 <= A[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-sightseeing-pair
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxScoreSightseeingPair(A []int) int {
	if len(A) == 0 {
		return 0
	}
	out := -1 << 31
	mx := A[0]
	for i := 1; i < len(A); i++ {
		out = Max(out, mx+A[i]-i)
		mx = Max(mx, A[i]+i)
	}
	return out
}
