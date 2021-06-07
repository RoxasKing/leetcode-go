package main

/*
  You are given an integer array nums and an integer target.

  You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

    For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".

  Return the number of different expressions that you can build, which evaluates to target.

  Example 1:
    Input: nums = [1,1,1,1,1], target = 3
    Output: 5
    Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
      -1 + 1 + 1 + 1 + 1 = 3
      +1 - 1 + 1 + 1 + 1 = 3
      +1 + 1 - 1 + 1 + 1 = 3
      +1 + 1 + 1 - 1 + 1 = 3
      +1 + 1 + 1 + 1 - 1 = 3

  Example 2:
    Input: nums = [1], target = 1
    Output: 1

  Constraints:
    1. 1 <= nums.length <= 20
    2. 0 <= nums[i] <= 1000
    3. 0 <= sum(nums[i]) <= 1000
    4. -1000 <= target <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/target-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
// Math

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if (sum+target)&1 == 1 {
		return 0
	}

	target = (sum + target) >> 1
	f := make([]int, target+1)
	f[0] = 1
	for _, num := range nums {
		for i := target; i >= num; i-- {
			f[i] += f[i-num]
		}
	}
	return f[target]
}
