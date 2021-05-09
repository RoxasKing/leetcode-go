package main

/*
  Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

  Example 1:
    Input: nums = [1,2,3]
    Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

  Example 2:
    Input: nums = [0,1]
    Output: [[0,1],[1,0]]

  Example 3:
    Input: nums = [1]
    Output: [[1]]

  Constraints:
    1. 1 <= nums.length <= 6
    2. -10 <= nums[i] <= 10
    3. All the integers of nums are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/permutations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func permute(nums []int) [][]int {
	out := [][]int{}
	dfs(nums, len(nums), 0, &out)
	return out
}

func dfs(nums []int, n, i int, out *[][]int) {
	if i == n {
		tmp := make([]int, n)
		copy(tmp, nums)
		*out = append(*out, tmp)
		return
	}

	for j := i; j < n; j++ {
		nums[i], nums[j] = nums[j], nums[i]
		dfs(nums, n, i+1, out)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
