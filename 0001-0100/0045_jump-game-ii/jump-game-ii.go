package main

func jump(nums []int) int {
	var steps, cur, max int
	for i := 0; i < len(nums)-1; i++ {
		max = Max(max, nums[i]+i)
		if i == cur {
			steps++
			cur = max
		}
	}
	return steps
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
