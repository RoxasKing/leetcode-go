package main

/*
  Given an array of integers numbers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

  Return the indices of the two numbers (1-indexed) as an integer array answer of size 2, where 1 <= answer[0] < answer[1] <= numbers.length.

  You may assume that each input would have exactly one solution and you may not use the same element twice.

  Example 1:
    Input: numbers = [2,7,11,15], target = 9
    Output: [1,2]
    Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.

  Example 2:
    Input: numbers = [2,3,4], target = 6
    Output: [1,3]

  Example 3:
    Input: numbers = [-1,0], target = -1
    Output: [1,2]

  Constraints:
    1. 2 <= numbers.length <= 3 * 10^4
    2. -1000 <= numbers[i] <= 1000
    3. numbers is sorted in increasing order.
    4. -1000 <= target <= 1000
    5. Only one valid answer exists.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	l, r := 0, n-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			break
		}
	}
	return []int{l + 1, r + 1}
}
