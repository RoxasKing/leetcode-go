package main

/*
  Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

  Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

  Example 1:
    Input: nums = [1,2,2,3,1]
    Output: 2
    Explanation:
    The input array has a degree of 2 because both elements 1 and 2 appear twice.
    Of the subarrays that have the same degree:
    [1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
    The shortest length is 2. So return 2.

  Example 2:
    Input: nums = [1,2,2,3,1,4,2]
    Output: 6
    Explanation:
    The degree is 3 because the element 2 is repeated 3 times.
    So [2,2,3,1,4,2] is the shortest subarray, therefore returning 6.

  Constraints:
    1. nums.length will be between 1 and 50,000.
    2. nums[i] will be an integer between 0 and 49,999.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/degree-of-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func findShortestSubArray(nums []int) int {
	n := len(nums)

	max := 0
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
		max = Max(max, cnt[num])
	}

	out := n
	cnt = map[int]int{}
	for l, r := 0, 0; r < n; r++ {
		cnt[nums[r]]++
		if cnt[nums[r]] < max {
			continue
		}
		for l <= r && nums[l] != nums[r] {
			cnt[nums[l]]--
			l++
		}
		out = Min(out, r+1-l)
	}

	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
