package main

/*
  Given an array of integers nums and an integer threshold, we will choose a positive integer divisor, divide all the array by it, and sum the division's result. Find the smallest divisor such that the result mentioned above is less than or equal to threshold.

  Each result of the division is rounded to the nearest integer greater than or equal to that element. (For example: 7/3 = 3 and 10/2 = 5).

  It is guaranteed that there will be an answer.

  Example 1:
    Input: nums = [1,2,5,9], threshold = 6
    Output: 5
    Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1.
      If the divisor is 4 we can get a sum of 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2).

  Example 2:
    Input: nums = [44,22,33,11,1], threshold = 5
    Output: 44

  Example 3:
    Input: nums = [21212,10101,12121], threshold = 1000000
    Output: 1

  Example 4:
    Input: nums = [2,3,5,7,11], threshold = 11
    Output: 3

  Constraints:
    1. 1 <= nums.length <= 5 * 10^4
    2. 1 <= nums[i] <= 10^6
    3. nums.length <= threshold <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-the-smallest-divisor-given-a-threshold
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, int(1e6)
	for l < r {
		m := (l + r) >> 1
		if getSum(nums, m) > threshold {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func getSum(nums []int, divisor int) int {
	out := 0
	for _, num := range nums {
		out += num / divisor
		if num%divisor > 0 {
			out++
		}
	}
	return out
}
