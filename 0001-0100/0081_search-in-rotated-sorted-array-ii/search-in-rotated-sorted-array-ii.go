package main

// Tags:
// Binary Search
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l < r && nums[l] == nums[r] {
		l++
	}
	rotateIdx := bSearchRotateIdx(nums, l, r)
	if rotateIdx == l || nums[l] > target {
		return bSearch(nums, rotateIdx, r, target)
	}
	return bSearch(nums, l, rotateIdx-1, target)
}

func bSearchRotateIdx(nums []int, l, r int) int {
	start := l
	for l < r {
		m := (l + r) >> 1
		if nums[m] > nums[m+1] {
			return m + 1
		}
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return start
}

func bSearch(nums []int, l, r, target int) bool {
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}
