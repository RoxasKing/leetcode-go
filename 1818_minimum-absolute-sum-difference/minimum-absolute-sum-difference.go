package main

import (
	"container/heap"
	"sort"
)

/*
  You are given two positive integer arrays nums1 and nums2, both of length n.

  The absolute sum difference of arrays nums1 and nums2 is defined as the sum of |nums1[i] - nums2[i]| for each 0 <= i < n (0-indexed).

  You can replace at most one element of nums1 with any other element in nums1 to minimize the absolute sum difference.

  Return the minimum absolute sum difference after replacing at most one element in the array nums1. Since the answer may be large, return it modulo 109 + 7.

  |x| is defined as:
    1. x if x >= 0, or
    2. -x if x < 0.

  Example 1:
    Input: nums1 = [1,7,5], nums2 = [2,3,5]
    Output: 3
    Explanation: There are two possible optimal solutions:
    - Replace the second element with the first: [1,7,5] => [1,1,5], or
    - Replace the second element with the third: [1,7,5] => [1,5,5].
    Both will yield an absolute sum difference of |1-2| + (|1-3| or |5-3|) + |5-5| = 3.

  Example 2:
    Input: nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
    Output: 0
    Explanation: nums1 is equal to nums2 so no replacement is needed. This will result in an
    absolute sum difference of 0.

  Example 3:
    Input: nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
    Output: 20
    Explanation: Replace the first element with the second: [1,10,4,4,2,7] => [10,10,4,4,2,7].
    This yields an absolute sum difference of |10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20

  Constraints:
    1. n == nums1.length
    2. n == nums2.length
    3. 1 <= n <= 10^5
    4. 1 <= nums1[i], nums2[i] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-absolute-sum-difference
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm + Binary Search + Priority Queue
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	out := 0
	pq := &PriorityQueue{}
	mark := make(map[int]bool)
	nums := []int{}

	for i := range nums1 {
		diff := Abs(nums1[i] - nums2[i])
		out = (out + diff) % (1e9 + 7)
		heap.Push(pq, [2]int{diff, i})
		if !mark[nums1[i]] {
			mark[nums1[i]] = true
			nums = append(nums, nums1[i])
		}
	}

	if out == 0 {
		return 0
	}

	sort.Ints(nums)

	max := 0
	for pq.Len() > 0 {
		e := heap.Pop(pq).([2]int)
		diff, j := e[0], e[1]
		if Abs(diff) <= max {
			break
		}

		i := sort.Search(len(nums), func(i int) bool { return nums[i] >= nums2[j] })

		if i >= 0 && i < len(nums) {
			max = Max(max, Abs(diff)-Abs(nums[i]-nums2[j]))
		}
		if i-1 >= 0 {
			max = Max(max, Abs(diff)-Abs(nums[i-1]-nums2[j]))
		}
		if i+1 < len(nums) {
			max = Max(max, Abs(diff)-Abs(nums[i+1]-nums2[j]))
		}
	}

	return out - max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i][0] > pq[j][0] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([2]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
