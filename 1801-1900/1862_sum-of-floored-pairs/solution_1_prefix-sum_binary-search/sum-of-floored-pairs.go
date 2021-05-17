package main

import "sort"

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

// Prefix Sum + Binary Search

func sumOfFlooredPairs(nums []int) int {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num]++
	}
	arr := make([]int, 0, len(freq))
	for num := range freq {
		arr = append(arr, num)
	}
	sort.Ints(arr)

	n := len(arr)
	pFreq := make([]int, n+1)
	for i := 0; i < n; i++ {
		pFreq[i+1] = pFreq[i] + freq[arr[i]]
	}

	out := 0
	for i, num := range arr {
		for base := num; i < n; base += num {
			j := sort.Search(n-i, func(j int) bool { return arr[j+i] >= base }) + i
			out += (pFreq[n] - pFreq[j]) * freq[num]
			i = j
		}
	}
	return out % (1e9 + 7)
}
