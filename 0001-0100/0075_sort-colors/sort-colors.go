package main

// Tags:
// Two Pointers

func sortColors(nums []int) {
	n := len(nums)
	l, m, r := 0, 0, n-1
	for m <= r {
		if nums[m] == 0 {
			nums[l], nums[m] = nums[m], nums[l]
			l++
			m++
		} else if nums[m] == 2 {
			nums[m], nums[r] = nums[r], nums[m]
			r--
		} else {
			m++
		}
	}
}
