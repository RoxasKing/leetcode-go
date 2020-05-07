package interview

/*
  找出数组中重复的数字。
  在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findRepeatNumber(nums []int) int {
	save := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := save[num]; ok {
			return num
		}
		save[num] = struct{}{}
	}
	return 0
}

func findRepeatNumber2(nums []int) int {
	for i := range nums {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			index := nums[i]
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return -1
}
