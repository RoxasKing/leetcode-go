package main

// Tags:
// Two Pointers
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}

// Binary Search
func findDuplicate2(nums []int) int {
	out := 0
	for l, r := 1, len(nums)-1; l <= r; {
		target := (l + r) >> 1
		cnt := 0
		for _, num := range nums {
			if num <= target {
				cnt++
			}
		}
		if cnt <= target {
			l = target + 1
		} else {
			r = target - 1
			out = target
		}
	}
	return out
}

// Bit Manipulation
func findDuplicate3(nums []int) int {
	out := 0
	maxBit := 31
	for (len(nums)-1)>>maxBit == 0 {
		maxBit--
	}
	for bit := 0; bit <= maxBit; bit++ {
		var x, y int
		for i, num := range nums {
			if num&(1<<bit) != 0 {
				x++
			}
			if i > 0 && i&(1<<bit) != 0 {
				y++
			}
		}
		if x > y {
			out |= 1 << bit
		}
	}
	return out
}
