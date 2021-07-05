package main

func checkPossibility(nums []int) bool {
	n := len(nums)
	changed := false
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			if changed {
				return false
			}
			if i-1 >= 0 {
				if nums[i+1] >= nums[i-1] {
					nums[i] = nums[i-1]
				} else {
					nums[i+1] = nums[i]
				}
			} else {
				nums[i] = nums[i+1]
			}
			changed = true
		}
	}
	return true
}
