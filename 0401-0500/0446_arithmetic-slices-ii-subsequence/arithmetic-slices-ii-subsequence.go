package main

// Tags:
// Dynamic Programming
// Hash

func numberOfArithmeticSlices(nums []int) int {
	var out int
	n := len(nums)
	f := make([]map[int]int, n)
	for i := range f {
		f[i] = map[int]int{}
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			sum := f[j][diff]
			f[i][diff] += sum + 1
			out += sum
		}
	}
	return out
}
