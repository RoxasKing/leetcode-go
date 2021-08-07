package main

// Tags:
// Two Pointers

func circularArrayLoop(nums []int) bool {
	n := len(nums)
	next := func(i int) int {
		return ((i+nums[i])%n + n) % n
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}

		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				if slow == next(fast) {
					break
				}
				return true
			}
			slow, fast = next(slow), next(next(fast))
		}
		ptr := i
		for nums[ptr]*nums[next(ptr)] > 0 {
			tmp := ptr
			ptr = next(ptr)
			nums[tmp] = 0
		}
	}
	return false
}
