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
		if nums[i]&1 == 1 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	return nums
}

func exchange2(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]&1 == 1 {
			l++
		}
		for l < r && nums[r]&1 == 0 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
