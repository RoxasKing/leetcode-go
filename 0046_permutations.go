package My_LeetCode_In_Go

/*
  给定一个没有重复数字的序列，返回其所有可能的全排列。
*/

func permute(nums []int) [][]int {
	var out [][]int
	var backtrack func(nums []int, index int)
	backtrack = func(nums []int, index int) {
		if index == len(nums) {
			out = append(out, nums)
		}
		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			backtrack(append([]int{}, nums...), index+1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	backtrack(append([]int{}, nums...), 0)
	return out
}
