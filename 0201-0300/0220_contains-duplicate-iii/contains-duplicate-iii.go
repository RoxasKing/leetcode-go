package main

import "sort"

/*
  Given an integer array nums and two integers k and t, return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.

  Example 1:
    Input: nums = [1,2,3,1], k = 3, t = 0
    Output: true

  Example 2:
    Input: nums = [1,0,1,1], k = 1, t = 2
    Output: true

  Example 3:
    Input: nums = [1,5,9,1,5,9], k = 2, t = 3
    Output: false

  Constraints:
    1. 0 <= nums.length <= 2 * 10^4
    2. -2^31 <= nums[i] <= 2^31 - 1
    3. 0 <= k <= 10^4
    4. 0 <= t <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/contains-duplicate-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bucket Sort
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	buckets := make(map[int]int)
	for i, a := range nums {
		bucket := getID(a, t+1)
		if _, ok := buckets[bucket]; ok {
			return true
		}
		if b, ok := buckets[bucket-1]; ok && Abs(a-b) <= t {
			return true
		}
		if b, ok := buckets[bucket+1]; ok && Abs(a-b) <= t {
			return true
		}
		buckets[bucket] = a
		if i >= k {
			delete(buckets, getID(nums[i-k], t+1))
		}
	}
	return false
}

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

// Binary Search
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	n := len(nums)
	idxs := make([]int, n)
	for i := range idxs {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return nums[idxs[i]] < nums[idxs[j]]
	})
	for i := range nums {
		l := sort.Search(n, func(j int) bool {
			return nums[idxs[j]] >= nums[i]-t
		})
		r := sort.Search(n, func(j int) bool {
			return nums[idxs[j]] > nums[i]+t
		}) - 1
		for j := l; j <= r; j++ {
			if idxs[j] == i {
				continue
			}
			if Abs(i-idxs[j]) <= k {
				return true
			}
		}
	}
	return false
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
