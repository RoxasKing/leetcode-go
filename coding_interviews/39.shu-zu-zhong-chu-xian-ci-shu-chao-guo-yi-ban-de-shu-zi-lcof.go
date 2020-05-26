package codinginterviews

/*
  数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
  你可以假设数组是非空的，并且给定的数组总是存在多数元素。
  限制：
    1 <= 数组长度 <= 50000
*/

// Boyer-Moore
func majorityElement(nums []int) int {
	var cur, count int
	for _, num := range nums {
		if count == 0 {
			cur = num
		}
		if num == cur {
			count++
		} else {
			count--
		}
	}
	return cur
}

// Divide and conquer
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
