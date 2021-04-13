package main

/*
  There is a function signFunc(x) that returns:
    1. 1 if x is positive.
    2. -1 if x is negative.
    3. 0 if x is equal to 0.
  You are given an integer array nums. Let product be the product of all values in the array nums.

  Return signFunc(product).

  Example 1:
    Input: nums = [-1,-2,-3,-4,3,2,1]
    Output: 1
    Explanation: The product of all values in the array is 144, and signFunc(144) = 1

  Example 2:
    Input: nums = [1,5,0,2,-3]
    Output: 0
    Explanation: The product of all values in the array is 0, and signFunc(0) = 0

  Example 3:
    Input: nums = [-1,1,-1,1,-1]
    Output: -1
    Explanation: The product of all values in the array is -1, and signFunc(-1) = -1

  Constraints:
    1. 1 <= nums.length <= 1000
    2. -100 <= nums[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sign-of-the-product-of-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func arraySign(nums []int) int {
	out := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			out = -out
		}
	}
	return out
}
