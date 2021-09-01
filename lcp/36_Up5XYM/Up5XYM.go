package main

import "sort"

// Tags:
// Dynamic Programming

func maxGroupNumber(tiles []int) int {
	freq := map[int]int{}
	nums := []int{}
	for _, tile := range tiles {
		if freq[tile] == 0 {
			nums = append(nums, tile)
		}
		freq[tile]++
	}
	sort.Ints(nums)

	reset := func(f [][]int) {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				f[i][j] = -1
			}
		}
	}

	f0 := make([][]int, 5)
	f1 := make([][]int, 5)
	for i := 0; i < 5; i++ {
		f0[i] = make([]int, 5)
		f1[i] = make([]int, 5)
	}
	reset(f0)
	f0[0][0] = 0

	pre := -1
	for _, num := range nums {
		if num != pre+1 {
			f00 := f0[0][0]
			reset(f0)
			f0[0][0] = f00
		}
		reset(f1)
		cnt0 := freq[num]
		for cnt2 := 0; cnt2 <= Min(4, cnt0); cnt2++ {
			for cnt1 := cnt2; cnt1 <= Min(4, cnt0); cnt1++ {
				if f0[cnt2][cnt1] < 0 {
					continue
				}
				for shunzi := 0; shunzi <= cnt2; shunzi++ {
					new2 := cnt1 - shunzi
					for new1 := 0; new1 <= Min(4, cnt0-shunzi); new1++ {
						f1[new2][new1] = Max(f1[new2][new1], f0[cnt2][cnt1]+shunzi+(cnt0-shunzi-new1)/3)
					}
				}
			}
		}
		f0, f1 = f1, f0
		pre = num
	}
	return f0[0][0]
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
