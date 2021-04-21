package main

/*
  Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

  Example 1:
    Input: nums = [1,2,3]
    Output: 6

  Example 2:
    Input: nums = [1,2,3,4]
    Output: 24

  Example 3:
    Input: nums = [-1,-2,-3]
    Output: -6

  Constraints:
    3 <= nums.length <= 10^4
    -1000 <= nums[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-product-of-three-numbers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximumProduct(nums []int) int {
	max0, max1, max2 := -1<<31, -1<<31, -1<<31
	min0, min1 := 1<<31-1, 1<<31-1
	for _, num := range nums {
		if num > max2 {
			max0, max1, max2 = max1, max2, num
		} else if num > max1 {
			max0, max1 = max1, num
		} else if num > max0 {
			max0 = num
		}

		if num < min0 {
			min0, min1 = num, min0
		} else if num < min1 {
			min1 = num
		}
	}

	if max2 < 0 {
		return max0 * max1 * max2
	} else {
		return Max(min0*min1, max0*max1) * max2
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
