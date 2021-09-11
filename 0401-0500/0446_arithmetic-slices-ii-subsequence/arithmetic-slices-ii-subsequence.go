package main

// Tags:
// Dynamic Programming
// Hash

func numberOfArithmeticSlices(nums []int) int {
	out := 0
	f := make([]map[int]int, len(nums))
	for i := range f {
		f[i] = map[int]int{}
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			c := f[j][d]
			f[i][d] += c + 1
			out += c
		}
	}
	return out
}
