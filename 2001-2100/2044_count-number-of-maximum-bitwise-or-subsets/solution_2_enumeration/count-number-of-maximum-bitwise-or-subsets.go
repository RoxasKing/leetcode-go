package main

// Difficulty:
// Medium

// Tags:
// Enumeration

func countMaxOrSubsets(nums []int) int {
	out, target := 0, 0
	for _, num := range nums {
		target |= num
	}
	n := len(nums)
	for mask := 1; mask <= 1<<n; mask++ {
		or := 0
		for i := n - 1; i >= 0; i-- {
			if (mask>>i)&1 == 1 {
				or |= nums[n-1-i]
			}
		}
		if or == target {
			out++
		}
	}
	return out
}
