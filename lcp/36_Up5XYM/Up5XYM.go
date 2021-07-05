package main

import "sort"

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
