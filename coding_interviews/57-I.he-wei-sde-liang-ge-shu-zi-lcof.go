package codinginterviews

/*
  输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

  限制：
    1 <= nums.length <= 10^5
    1 <= nums[i] <= 10^6
*/

func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		switch {
		case sum < target:
			l++
		case sum > target:
			r--
		default:
			return []int{nums[l], nums[r]}
		}
	}
	return nil
}
