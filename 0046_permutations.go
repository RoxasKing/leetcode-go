package leetcode

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

func permute2(nums []int) [][]int {
	var out [][]int
	permuteBacktrack(nums, 0, &out)
	return out
}

func permuteBacktrack(nums []int, index int, out *[][]int) {
	if index == len(nums) {
		*out = append(*out, nums)
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		permuteBacktrack(newNums, index+1, out)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
