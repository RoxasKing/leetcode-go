package codinginterviews

/*
  输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
  提示：
    1 <= nums.length <= 50000
    1 <= nums[i] <= 10000
*/

func exchange(nums []int) []int {
	var index int
	for i := range nums {
		if isOdd(nums[i]) {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	return nums
}

func isOdd(num int) bool {
	return num&1 == 1
}

func exchange2(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && isOdd(nums[l]) {
			l++
		}
		for l < r && !isOdd(nums[r]) {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
