package main

import (
	"container/heap"
	"sort"
)

/*
  Given the array nums consisting of n positive integers. You computed the sum of all non-empty continous subarrays from the array and then sort them in non-decreasing order, creating a new array of n * (n + 1) / 2 numbers.

  Return the sum of the numbers from index left to index right (indexed from 1), inclusive, in the new array. Since the answer can be a huge number return it modulo 10^9 + 7.

  Example 1:
    Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
    Output: 13
    Explanation: All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. After sorting them in non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13.

  Example 2:
    Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
    Output: 6
    Explanation: The given array is the same as example 1. We have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. The sum of the numbers from index le = 3 to ri = 4 is 3 + 3 = 6.

  Example 3:
    Input: nums = [1,2,3,4], n = 4, left = 1, right = 10
    Output: 50

  Constraints:
    1. 1 <= nums.length <= 10^3
    2. nums.length == n
    3. 1 <= nums[i] <= 100
    4. 1 <= left <= right <= n * (n + 1) / 2

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/range-sum-of-sorted-subarray-sums
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Prefix Sum + Binary Search + Two Pointers
func rangeSum(nums []int, n int, left int, right int) int {
	pSum := make([]int, n+1)
	ppSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
		ppSum[i+1] = ppSum[i] + pSum[i+1]
	}
	mod := int(1e9 + 7)
	return (theKthSum(pSum, ppSum, n, right) - theKthSum(pSum, ppSum, n, left-1) + mod) % mod
}

func theKthSum(pSum, ppSum []int, n, k int) int {
	num := bSearchTheKSum(pSum, n, k)
	sum, cnt := 0, 0
	for i, j := 0, 0; i <= n; i++ {
		for j+1 <= n && pSum[j+1]-pSum[i] < num {
			j++
		}
		sum += ppSum[j] - ppSum[i] - pSum[i]*(j-i)
		cnt += j - i
	}
	return sum + (k-cnt)*num
}

func bSearchTheKSum(pSum []int, n, k int) int {
	l, r := 0, pSum[n]
	for l < r {
		sum := (l + r) >> 1
		cnt := cntNoBiggerSum(pSum, n, sum)
		if cnt < k {
			l = sum + 1
		} else {
			r = sum
		}
	}
	return l
}

func cntNoBiggerSum(pSum []int, n, sum int) int {
	out := 0
	for i, j := 0, 0; i <= n; i++ {
		for j+1 <= n && pSum[j+1]-pSum[i] <= sum {
			j++
		}
		out += j - i
	}
	return out
}

// Prefix Sum + Priority Queue + Merge Sort
func rangeSum2(nums []int, n int, left int, right int) int {
	pq := PriorityQueue{}
	for i := 0; i < n; i++ {
		heap.Push(&pq, &prefixSum{sum: nums[i], idx: i})
	}
	idx := 0
	sum := 0
	for idx < right {
		idx++
		ps := heap.Pop(&pq).(*prefixSum)
		if idx >= left {
			sum += ps.sum
			sum %= 1e9 + 7
		}
		ps.idx++
		if ps.idx == n {
			continue
		}
		ps.sum += nums[ps.idx]
		heap.Push(&pq, ps)
	}
	return sum
}

type prefixSum struct {
	sum int
	idx int
}

type PriorityQueue []*prefixSum

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].sum < pq[j].sum }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*prefixSum)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := pq.Len() - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}

// Prefix Sum
func rangeSum3(nums []int, n int, left int, right int) int {
	pSum := make([]int, n)
	sums := make([]int, 0, n*(n+1)/2)
	for i := 0; i < n; i++ {
		pSum[i] = nums[i]
		if i > 0 {
			pSum[i] += pSum[i-1]
		}
		sums = append(sums, pSum[i])
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sums = append(sums, pSum[j]-pSum[i])
		}
	}
	sort.Ints(sums)
	sum := 0
	for i := left - 1; i < right; i++ {
		sum += sums[i]
		sum %= 1e9 + 7
	}
	return sum
}
