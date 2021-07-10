package main

// Tags:
// Boyer-Moore

func majorityElement(nums []int) int {
	var cur, cnt int
	for _, num := range nums {
		if cnt == 0 {
			cur = num
		}
		if num == cur {
			cnt++
		} else {
			cnt--
		}
	}
	cnt = 0
	for _, num := range nums {
		if num == cur {
			cnt++
		}
	}
	if cnt > len(nums)>>1 {
		return cur
	}
	return -1
}
