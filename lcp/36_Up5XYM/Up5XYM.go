package main

import "sort"

/*
  麻将的游戏规则中，共有两种方式凑成「一组牌」：
    1. 顺子：三张牌面数字连续的麻将，例如 [4,5,6]
    2. 刻子：三张牌面数字相同的麻将，例如 [10,10,10]
  给定若干数字作为麻将牌的数值（记作一维数组 tiles），请返回所给 tiles 最多可组成的牌组数。

  注意：凑成牌组时，每张牌仅能使用一次。

  示例 1：
    输入：tiles = [2,2,2,3,4]
    输出：1
    解释：最多可以组合出 [2,2,2] 或者 [2,3,4] 其中一组牌。

  示例 2：
    输入：tiles = [2,2,2,3,4,1,3]
    输出：2
    解释：最多可以组合出 [1,2,3] 与 [2,3,4] 两组牌。

  提示：
    1. 1 <= tiles.length <= 10^5
    2. 1 <= tiles[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/Up5XYM
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func maxGroupNumber(tiles []int) int {
	titleCnt := make(map[int]int)
	titleArr := []int{}
	for _, title := range tiles {
		if titleCnt[title] == 0 {
			titleArr = append(titleArr, title)
		}
		titleCnt[title]++
	}
	sort.Ints(titleArr)

	dp := make([][5][5]int, 2)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			dp[0][i][j] = -1
		}
	}
	dp[0][0][0] = 0

	prevTitle := -1
	for _, title := range titleArr {
		cnt := titleCnt[title]
		if prevTitle != title-1 {
			dp00 := dp[0][0][0]
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					dp[0][i][j] = -1
				}
			}
			dp[0][0][0] = dp00
		}

		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				dp[1][i][j] = -1
			}
		}

		for cnt2 := 0; cnt2 < 5; cnt2++ {
			for cnt1 := 0; cnt1 < 5; cnt1++ {
				if dp[0][cnt2][cnt1] < 0 {
					continue
				}

				for shunzi := 0; shunzi <= Min(cnt2, Min(cnt1, cnt)); shunzi++ {
					new2 := cnt1 - shunzi

					for new1 := 0; new1 <= Min(4, cnt-shunzi); new1++ {
						newScore := dp[0][cnt2][cnt1] + shunzi + (cnt-shunzi-new1)/3
						dp[1][new2][new1] = Max(dp[1][new2][new1], newScore)
					}
				}
			}
		}

		dp[0], dp[1] = dp[1], dp[0]
		prevTitle = title
	}

	out := 0
	for cnt2 := 0; cnt2 < 5; cnt2++ {
		for cnt1 := 0; cnt1 < 5; cnt1++ {
			out = Max(out, dp[0][cnt2][cnt1])
		}
	}
	return out
}

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
