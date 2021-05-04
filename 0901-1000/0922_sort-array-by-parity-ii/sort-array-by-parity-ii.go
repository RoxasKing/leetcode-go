package main

/*
  Given an array of integers nums, half of the integers in nums are odd, and the other half are even.

  Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.

  Return any answer array that satisfies this condition.

  Example 1:
    Input: nums = [4,2,5,7]
    Output: [4,5,2,7]
    Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.

  Example 2:
    Input: nums = [2,3]
    Output: [2,3]

  Constraints:
    1. 2 <= nums.length <= 2 * 10^4
    2. nums.length is even.
    3. Half of the integers in nums are even.
    4. 0 <= nums[i] <= 1000

  Follow Up: Could you solve it in-place?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sort-array-by-parity-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Two Pointers

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	for i, j := 0, 1; i < n && j < n; {
		if nums[i]&1 == 1 && nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i += 2
			j += 2
			continue
		}
		if nums[i]&1 == 0 {
			i += 2
		}
		if nums[j]&1 == 1 {
			j += 2
		}
	}
	return nums
}
