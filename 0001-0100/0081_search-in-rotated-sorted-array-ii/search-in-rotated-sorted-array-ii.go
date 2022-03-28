package main

// Difficulty:
// Medium

// Tags:
// Binary Search

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for ; l < r && nums[l] == nums[r]; l++ {
	}
	rotateIdx := l
	for i, j := l, r; i < j; {
		k := (i + j) >> 1
		if nums[k] > nums[k+1] {
			rotateIdx = k + 1
			break
		}
		if nums[i] <= nums[k] { // Attention!! (e.g: [2,2,2,2,2,3,4,1], nums[i] == nums[k] == 2)
			i = k + 1
		} else {
			j = k
		}
	}
	if nums[l] > target {
		l = rotateIdx
	} else if rotateIdx != l {
		r = rotateIdx - 1
	}
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			return true
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
