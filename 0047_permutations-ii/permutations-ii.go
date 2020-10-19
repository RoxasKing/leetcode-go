package main

import (
	"sort"
)

/*
  给定一个可包含重复数字的序列，返回所有不重复的全排列。

  示例:
    输入: [1,1,2]
    输出:
    [
      [1,1,2],
      [1,2,1],
      [2,1,1]
    ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/permutations-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func permuteUnique(nums []int) [][]int {
	var out [][]int
	n := len(nums)
	sort.Ints(nums)
	used := make([]bool, n)
	cur := make([]int, 0, n)
	backtrack(nums, n, &out, used, &cur, 0)
	return out
}

func backtrack(nums []int, n int, out *[][]int, used []bool, cur *[]int, index int) {
	if index == n {
		tmp := make([]int, n)
		copy(tmp, *cur)
		*out = append(*out, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		*cur = append(*cur, nums[i])
		used[i] = true
		backtrack(nums, n, out, used, cur, index+1)
		used[i] = false
		*cur = (*cur)[:len(*cur)-1]
	}
}
