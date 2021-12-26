package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	target := sum >> 1
	f := make([]bool, target+1)
	f[0] = true
	for _, num := range nums {
		for i := target; i >= num; i-- {
			f[i] = f[i] || f[i-num]
		}
	}
	return f[target]
}
