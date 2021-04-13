package main

/*
  给定一个 没有重复 数字的序列，返回其所有可能的全排列。

  示例:
    输入: [1,2,3]
    输出:
    [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/permutations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func permute(nums []int) [][]int {
	var out [][]int
	backtrack(nums, &out, 0)
	return out
}

func backtrack(nums []int, out *[][]int, index int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*out = append(*out, tmp)
		return
	}
	for i := index; i < len(nums); i++ {
		nums[i], nums[index] = nums[index], nums[i]
		backtrack(nums, out, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
