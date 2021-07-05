package main

// Tags:
// Binary Search
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	start := -1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == 0 || nums[m-1] < target {
				start = m
				break
			} else {
				r = m - 1
			}
		}
	}
	if start == -1 {
		return 0
	}

	l, r = start, n-1
	end := -1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == n-1 || nums[m+1] > target {
				end = m
				break
			} else {
				l = m + 1
			}
		}
	}
	return end + 1 - start
}
