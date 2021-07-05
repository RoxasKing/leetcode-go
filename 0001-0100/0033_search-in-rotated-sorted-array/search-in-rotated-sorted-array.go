package main

// Tags:
// Binary Search
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
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
		if nums[l] < nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return start
}

func bSearch(nums []int, l, r, target int) int {
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
