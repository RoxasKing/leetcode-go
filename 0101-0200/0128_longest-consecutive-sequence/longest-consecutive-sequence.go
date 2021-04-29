package main

/*
  Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

  Example 1:
    Input: nums = [100,4,200,1,3,2]
    Output: 4
    Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

  Example 2:
    Input: nums = [0,3,7,2,5,8,4,6,0,1]
    Output: 9

  Constraints:
    1. 0 <= nums.length <= 10^4
    2. -10^9 <= nums[i] <= 10^9

  Follow up: Could you implement the O(n) solution?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash
func longestConsecutive(nums []int) int {
	contain := make(map[int]bool)
	for _, num := range nums {
		contain[num] = true
	}
	out := 0
	for _, num := range nums {
		if contain[num-1] {
			continue
		}
		cnt := 1
		for i := num + 1; contain[i]; i++ {
			cnt++
		}
		out = Max(out, cnt)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
