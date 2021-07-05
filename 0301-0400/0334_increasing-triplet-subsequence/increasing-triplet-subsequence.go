package main

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	i, j, k := len(nums)-3, len(nums)-2, len(nums)-1
	for i >= 0 {
		if nums[i] < nums[j] && nums[j] < nums[k] {
			return true
		}
		if nums[i+1] > nums[k] {
			k = i + 1
			j = k - 1
		} else if nums[i] < nums[k] && nums[i] > nums[j] {
			j = i
		}
		i--
	}
	return false
}

// Monotone Stack
func increasingTriplet2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var stack []int
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
		if len(stack) == 3 {
			return true
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
		if len(stack) == 3 {
			return true
		}
	}
	return false
}
