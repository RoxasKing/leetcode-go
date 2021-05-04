package main

/*
  Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

  Return true if there is a 132 pattern in nums, otherwise, return false.

  Follow up: The O(n^2) is trivial, could you come up with the O(n logn) or the O(n) solution?

  Example 1:
    Input: nums = [1,2,3,4]
    Output: false
    Explanation: There is no 132 pattern in the sequence.

  Example 2:
    Input: nums = [3,1,4,2]
    Output: true
    Explanation: There is a 132 pattern in the sequence: [1, 4, 2].

  Example 3:
    Input: nums = [-1,3,2,0]
    Output: true
    Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3, 0] and [-1, 2, 0].

  Constraints:
    1. n == nums.length
    2. 1 <= n <= 10^4
    3. -10^9 <= nums[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/132-pattern
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Monotone Stack
func find132pattern(nums []int) bool {
	n := len(nums)
	stk := []int{}
	num2 := -1 << 31
	for i := n - 1; i >= 0; i-- {
		if nums[i] < num2 {
			return true
		}
		for len(stk) > 0 && stk[len(stk)-1] < nums[i] {
			num2 = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		if nums[i] > num2 {
			stk = append(stk, nums[i])
		}
	}
	return false
}
