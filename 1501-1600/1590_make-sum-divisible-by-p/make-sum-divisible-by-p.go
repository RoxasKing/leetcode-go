package main

/*
  Given an array of positive integers nums, remove the smallest subarray (possibly empty) such that the sum of the remaining elements is divisible by p. It is not allowed to remove the whole array.

  Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.

  A subarray is defined as a contiguous block of elements in the array.

  Example 1:
    Input: nums = [3,1,4,2], p = 6
    Output: 1
    Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.

  Example 2:
    Input: nums = [6,3,5,2], p = 9
    Output: 2
    Explanation: We cannot remove a single element to get a sum divisible by 9. The best way is to remove the subarray [5,2], leaving us with [6,3] with sum 9.

  Example 3:
    Input: nums = [1,2,3], p = 3
    Output: 0
    Explanation: Here the sum is 6. which is already divisible by 3. Thus we do not need to remove anything.

  Example 4:
    Input: nums = [1,2,3], p = 7
    Output: -1
    Explanation: There is no way to remove a subarray in order to get a sum divisible by 7.

  Example 5:
    Input: nums = [1000000000,1000000000,1000000000], p = 3
    Output: 0

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 1 <= nums[i] <= 10^9
    3. 1 <= p <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/make-sum-divisible-by-p
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum + Hash

func minSubarray(nums []int, p int) int {
	out := -1
	n := len(nums)
	sSum := make([]int, n)

	for i, sum := n-1, 0; i >= 0; i-- {
		sum += nums[i]
		sum %= p
		if sum == 0 && (out == -1 || i < out) {
			out = i
		}
		sSum[i] = sum
	}

	dict := map[int]int{0: -1}

	for i, sum := 0, 0; i < n; i++ {
		sum += nums[i]
		sum %= p
		if sum == 0 && (out == -1 || n-(i+1) < out) {
			out = n - (i + 1)
		}
		if j, ok := dict[p-sSum[i]]; ok && (out == -1 || i-j-1 < out) {
			out = i - j - 1
		}
		dict[sum] = i
	}

	return out
}
