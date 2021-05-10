package main

/*
  The min-product of an array is equal to the minimum value in the array multiplied by the array's sum.

    For example, the array [3,2,5] (minimum value is 2) has a min-product of 2 * (3+2+5) = 2 * 10 = 20.

  Given an array of integers nums, return the maximum min-product of any non-empty subarray of nums. Since the answer may be large, return it modulo 109 + 7.

  Note that the min-product should be maximized before performing the modulo operation. Testcases are generated such that the maximum min-product without modulo will fit in a 64-bit signed integer.

  A subarray is a contiguous part of an array.

  Example 1:
    Input: nums = [1,2,3,2]
    Output: 14
    Explanation: The maximum min-product is achieved with the subarray [2,3,2] (minimum value is 2).
      2 * (2+3+2) = 2 * 7 = 14.

  Example 2:
    Input: nums = [2,3,3,1,2]
    Output: 18
    Explanation: The maximum min-product is achieved with the subarray [3,3] (minimum value is 3).
      3 * (3+3) = 3 * 6 = 18.

  Example 3:
    Input: nums = [3,1,5,6,4,2]
    Output: 60
    Explanation: The maximum min-product is achieved with the subarray [5,6,4] (minimum value is 4).
      4 * (5+6+4) = 4 * 15 = 60.

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= nums[i] <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-subarray-min-product
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum + Monotone Stack
func maxSumMinProduct(nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)
	for i, num := range nums {
		pSum[i+1] = pSum[i] + num
	}

	stk := IntStack{-1}
	out := 0
	for i, num := range nums {
		for stk.Len() > 1 && num <= nums[stk.Top()] {
			out = Max(out, nums[stk.Pop()]*(pSum[i]-pSum[stk.Top()+1]))
		}
		stk.Push(i)
	}

	for stk.Len() > 1 {
		out = Max(out, nums[stk.Pop()]*(pSum[n]-pSum[stk.Top()+1]))
	}

	return out % (1e9 + 7)
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
