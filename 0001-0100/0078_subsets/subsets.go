package main

/*
  Given an integer array nums of unique elements, return all possible subsets (the power set).

  The solution set must not contain duplicate subsets. Return the solution in any order.

  Example 1:
    Input: nums = [1,2,3]
    Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

  Example 2:
    Input: nums = [0]
    Output: [[],[0]]

  Constraints:
    1. 1 <= nums.length <= 10
    2. -10 <= nums[i] <= 10
    3. All the numbers of nums are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/subsets
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func subsets(nums []int) [][]int {
	out := [][]int{}
	dfs(nums, len(nums), 0, []int{}, &out)
	return out
}

func dfs(nums []int, n, i int, cur []int, out *[][]int) {
	if i == n {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*out = append(*out, tmp)
		return
	}

	cur = append(cur, nums[i])
	dfs(nums, n, i+1, cur, out)
	cur = cur[:len(cur)-1]

	dfs(nums, n, i+1, cur, out)
}
