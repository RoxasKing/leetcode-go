package main

/*
  You are given an integer array nums and two integers limit and goal. The array nums has an interesting property that abs(nums[i]) <= limit.

  Return the minimum number of elements you need to add to make the sum of the array equal to goal. The array must maintain its property that abs(nums[i]) <= limit.

  Note that abs(x) equals x if x >= 0, and -x otherwise.

  Example 1:
    Input: nums = [1,-1,1], limit = 3, goal = -4
    Output: 2
    Explanation: You can add -2 and -3, then the sum of the array will be 1 - 1 + 1 - 2 - 3 = -4.

  Example 2:
    Input: nums = [1,-10,9,1], limit = 100, goal = 0
    Output: 1

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= limit <= 10^6
    3. -limit <= nums[i] <= limit
    4. -10^9 <= goal <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-elements-to-add-to-form-a-given-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	dif := goal - sum
	out := Abs(dif) / Abs(limit)
	if Abs(dif)%Abs(limit) > 0 {
		out++
	}
	return out
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
