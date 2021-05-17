package main

/*
  Given an integer array nums, return the sum of floor(nums[i] / nums[j]) for all pairs of indices 0 <= i, j < nums.length in the array. Since the answer may be too large, return it modulo 109 + 7.

  The floor() function returns the integer part of the division.

  Example 1:
    Input: nums = [2,5,9]
    Output: 10
    Explanation:
      floor(2 / 5) = floor(2 / 9) = floor(5 / 9) = 0
      floor(2 / 2) = floor(5 / 5) = floor(9 / 9) = 1
      floor(5 / 2) = 2
      floor(9 / 2) = 4
      floor(9 / 5) = 1
      We calculate the floor of the division for every pair of indices in the array then sum them up.

  Example 2:
    Input: nums = [7,7,7,7,7,7,7]
    Output: 49

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= nums[i] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-floored-pairs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum + Math

func sumOfFlooredPairs(nums []int) int {
	upper := 0
	freq := map[int]int{}
	for _, num := range nums {
		upper = Max(upper, num)
		freq[num]++
	}

	pFreq := make([]int, upper+1)
	for i := 1; i <= upper; i++ {
		pFreq[i] = pFreq[i-1] + freq[i]
	}

	out := 0
	for y := 1; y <= upper; y++ {
		if freq[y] == 0 {
			continue
		}
		for d := 1; d*y <= upper; d++ {
			out += freq[y] * d * (pFreq[Min((d+1)*y-1, upper)] - pFreq[d*y-1])
		}
	}
	return out % (1e9 + 7)
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
