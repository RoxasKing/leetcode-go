package leetcode

/*
  给定一个可包含重复数字的序列，返回所有不重复的全排列。
*/

func permuteUnique(nums []int) [][]int {
	var (
		out       [][]int
		track     []int
		used      = make([]bool, len(nums))
		contain   func([]int, int) bool
		backtrack func()
	)
	contain = func(nums []int, num int) bool {
		for i := range nums {
			if used[i] {
				continue
			}
			if nums[i] == num {
				return true
			}
		}
		return false
	}
	backtrack = func() {
		if len(track) == len(nums) {
			out = append(out, append([]int{}, track...))
			return
		}
		for i := range nums {
			if used[i] || contain(nums[:i], nums[i]) {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			backtrack()
			used[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack()
	return out
}
