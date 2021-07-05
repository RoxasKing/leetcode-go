package main

// Tags:
// Boyer-Moore
func majorityElement(nums []int) int {
	cur, cnt := nums[0], 1
	for _, num := range nums[1:] {
		if num == cur {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			cur, cnt = num, 1
		}
	}
	return cur
}
