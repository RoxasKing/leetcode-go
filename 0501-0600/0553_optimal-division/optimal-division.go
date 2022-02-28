package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func optimalDivision(nums []int) string {
	n := len(nums)
	type state struct {
		val float64
		str string
	}
	f_min := make([][]state, n)
	f_max := make([][]state, n)
	for i := 0; i < n; i++ {
		f_min[i] = make([]state, n)
		f_max[i] = make([]state, n)
		for j := 0; j < n; j++ {
			f_min[i][j].val = 1<<31 - 1
			f_max[i][j].val = -1 << 31
		}
		f_min[i][i] = state{val: float64(nums[i]), str: strconv.Itoa(nums[i])}
		f_max[i][i] = state{val: float64(nums[i]), str: strconv.Itoa(nums[i])}
	}
	for k := 2; k <= n; k++ {
		for i := 0; i+k-1 < n; i++ {
			for j := i; j < i+k-1; j++ {
				if res := f_min[i][j].val / f_max[j+1][i+k-1].val; f_min[i][i+k-1].val > res {
					f_min[i][i+k-1].val = res
					if j+1 < i+k-1 {
						f_min[i][i+k-1].str = f_min[i][j].str + "/(" + f_max[j+1][i+k-1].str + ")"
					} else {
						f_min[i][i+k-1].str = f_min[i][j].str + "/" + f_max[j+1][i+k-1].str
					}

				}
				if res := f_max[i][j].val / f_min[j+1][i+k-1].val; f_max[i][i+k-1].val < res {
					f_max[i][i+k-1].val = res
					if j+1 < i+k-1 {
						f_max[i][i+k-1].str = f_max[i][j].str + "/(" + f_min[j+1][i+k-1].str + ")"
					} else {
						f_max[i][i+k-1].str = f_max[i][j].str + "/" + f_min[j+1][i+k-1].str
					}
				}
			}
		}
	}
	return f_max[0][n-1].str
}
