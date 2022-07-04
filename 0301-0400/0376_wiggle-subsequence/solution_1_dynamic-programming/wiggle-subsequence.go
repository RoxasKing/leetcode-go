package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, 2)
		f[i][0], f[i][1] = 1, 1
	}
	o := 1
	for i := 1; i < n; i++ {
		for k := 0; k < i; k++ {
			if nums[k] > nums[i] && f[i][0] < f[k][1]+1 {
				if f[i][0] = f[k][1] + 1; o < f[i][0] {
					o = f[i][0]
				}
			}
			if nums[k] < nums[i] && f[i][1] < f[k][0]+1 {
				if f[i][1] = f[k][0] + 1; o < f[i][1] {
					o = f[i][1]
				}
			}
		}
	}
	return o
}
