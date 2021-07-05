package main

// Tags:
// Dynamic Programming
// Math

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)&1 == 1 {
		return 0
	}

	target = (sum + target) >> 1
	f := make([]int, target+1)
	f[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			f[i] += f[i-num]
		}
	}
	return f[target]
}
