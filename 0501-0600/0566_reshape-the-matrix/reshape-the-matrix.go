package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}

	out := make([][]int, r)
	for i := range out {
		out[i] = make([]int, c)
	}

	i, j := 0, 0
	for _, arr := range nums {
		for _, num := range arr {
			out[i][j] = num
			j++
			if j == c {
				i++
				j = 0
			}
		}
	}
	return out
}
