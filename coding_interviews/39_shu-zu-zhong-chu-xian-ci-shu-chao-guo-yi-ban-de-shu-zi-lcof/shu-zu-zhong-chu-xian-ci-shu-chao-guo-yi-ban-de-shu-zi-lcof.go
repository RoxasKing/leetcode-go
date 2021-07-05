package main

// Tags:
// Boyer-Moore
func majorityElement(nums []int) int {
	var out, cnt int
	for _, num := range nums {
		if cnt == 0 {
			out = num
		}
		if num == out {
			cnt++
		} else {
			cnt--
		}
	}
	return out
}

// Divide-and-conquer
func majorityElement2(nums []int) int {
	count := func(l, r, num int) int {
		var count int
		for i := l; i <= r; i++ {
			if nums[i] == num {
				count++
			}
		}
		return count
	}
	var recur func(int, int) int
	recur = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		m := l + (r-l)>>1
		left := recur(l, m)
		right := recur(m+1, r)
		if left == right {
			return left
		}
		lCount := count(l, r, left)
		rCount := count(l, r, right)
		if lCount > rCount {
			return left
		}
		return right
	}
	return recur(0, len(nums)-1)
}
