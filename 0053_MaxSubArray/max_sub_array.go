package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := 0
	max_negative := -9999999
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max_negative && nums[i] <= 0 {
			max_negative = nums[i]
		}
		sum += nums[i]
		if sum < 0 {
			sum = 0
		}
		if sum > max {
			max = sum
		}
	}
	if max == 0 && max_negative != -9999999 {
		return max_negative
	}
	return max
}

func maxSubArray2(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	for _, n := range nums {
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
