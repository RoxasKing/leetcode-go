package main

// Difficulty:
// Medium

// Tags:
// Binary Search

func searchRange(nums []int, target int) []int {
	n := len(nums)
	head, tail := -1, -1
	for l, r := 0, n-1; l <= r; {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == 0 || nums[m-1] < target {
				head = m
				break
			}
			r = m - 1
		}
	}
	if head == -1 {
		return []int{head, tail}
	}
	for l, r := head, n-1; l <= r; {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == n-1 || target < nums[m+1] {
				tail = m
				break
			}
			l = m + 1
		}
	}
	return []int{head, tail}
}
