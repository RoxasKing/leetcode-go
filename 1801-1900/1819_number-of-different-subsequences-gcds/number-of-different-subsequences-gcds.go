package main

/*
  You are given an array nums that consists of positive integers.

  The GCD of a sequence of numbers is defined as the greatest integer that divides all the numbers in the sequence evenly.
    For example, the GCD of the sequence [4,6,16] is 2.

  A subsequence of an array is a sequence that can be formed by removing some elements (possibly none) of the array.
    For example, [2,5,10] is a subsequence of [1,2,1,2,4,1,5,10].

  Return the number of different GCDs among all non-empty subsequences of nums.

  Example 1:
    Input: nums = [6,10,3]
    Output: 5
    Explanation: The figure shows all the non-empty subsequences and their GCDs.
    The different GCDs are 6, 10, 3, 2, and 1.

  Example 2:
    Input: nums = [5,15,40,5,6]
    Output: 7

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= nums[i] <= 2 * 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-different-subsequences-gcds
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func countDifferentSubsequenceGCDs(nums []int) int {
	out := 0
	dict := make(map[int]bool)
	for _, num := range nums {
		if !dict[num] {
			out++
		}
		dict[num] = true
	}
	for i := 1; i <= 1e5; i++ {
		if dict[i] {
			continue
		}
		gcd := -1
		for j := i << 1; j <= 2*1e5; j += i {
			if !dict[j] {
				continue
			}
			if gcd != -1 {
				gcd = Gcd(gcd, j)
			} else {
				gcd = j
			}
			if gcd == i {
				out++
				break
			}
		}
	}
	return out
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
