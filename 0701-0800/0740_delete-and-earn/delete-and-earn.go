package main

/*
  Given an array nums of integers, you can perform operations on the array.

  In each operation, you pick any nums[i] and delete it to earn nums[i] points. After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.

  You start with 0 points. Return the maximum number of points you can earn by applying such operations.

  Example 1:
    Input: nums = [3,4,2]
    Output: 6
    Explanation: Delete 4 to earn 4 points, consequently 3 is also deleted.
      Then, delete 2 to earn 2 points.
      6 total points are earned.

  Example 2:
    Input: nums = [2,2,3,3,3,4]
    Output: 9
    Explanation: Delete 3 to earn 3 points, deleting both 2's and the 4.
      Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
      9 total points are earned.

  Constraints:
    1. 1 <= nums.length <= 2 * 10^4
    2. 1 <= nums[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/delete-and-earn
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func deleteAndEarn(nums []int) int {
	cnt := [1e4 + 1]int{}
	for _, num := range nums {
		cnt[num] += num
	}

	f := [1e4 + 1]int{}
	f[1] = cnt[1]
	for i := 2; i <= 1e4; i++ {
		f[i] = Max(f[i-1], f[i-2]+cnt[i])
	}
	return f[1e4]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
