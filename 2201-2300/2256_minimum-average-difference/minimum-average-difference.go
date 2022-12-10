package main

// Difficulty:
// Medium

// Tags:
// Prefix Sum

func minimumAverageDifference(nums []int) int {
	n := len(nums)
	rsum := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		rsum[i] = rsum[i+1] + nums[i]
	}
	o, min, sum := n-1, rsum[0]/n, 0
	for i, x := range nums[:n-1] {
		sum += x
		if dif := abs(sum/(i+1) - rsum[i+1]/(n-1-i)); min > dif || min == dif && o > i {
			o, min = i, dif
		}
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
