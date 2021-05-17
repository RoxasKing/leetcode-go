package main

import "sort"

/*
  Given an integer array arr, remove a subarray (can be empty) from arr such that the remaining elements in arr are non-decreasing.

  A subarray is a contiguous subsequence of the array.

  Return the length of the shortest subarray to remove.

  Example 1:
    Input: arr = [1,2,3,10,4,2,3,5]
    Output: 3
    Explanation: The shortest subarray we can remove is [10,4,2] of length 3. The remaining elements after that will be [1,2,3,3,5] which are sorted.
      Another correct solution is to remove the subarray [3,10,4].

  Example 2:
    Input: arr = [5,4,3,2,1]
    Output: 4
    Explanation: Since the array is strictly decreasing, we can only keep a single element. Therefore we need to remove a subarray of length 4, either [5,4,3,2] or [4,3,2,1].

  Example 3:
    Input: arr = [1,2,3]
    Output: 0
    Explanation: The array is already non-decreasing. We do not need to remove any elements.

  Example 4:
    Input: arr = [1]
    Output: 0

  Constraints:
    1. 1 <= arr.length <= 10^5
    2. 0 <= arr[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Binary Search

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	l, r := 0, n-1
	for l+1 <= r && arr[l] <= arr[l+1] { // 找出非递减的最大前缀 [0, l]
		l++
	}
	for l <= r-1 && arr[r-1] <= arr[r] { // 找出非递减的最大后缀 [r, n-1]
		r--
	}

	if l == r {
		return 0
	}

	out := Min(n-(l+1), r) // 比较删除 [l+1, n-1] 或 [0, r-1]，哪个删除的最少

	for i := l; i >= 0; i-- { // 当前缀为 [0, i] 时， 找出符合非递减的后缀 [j, n-1]
		j := sort.Search(n-r, func(j int) bool { return arr[j+r] >= arr[i] }) + r
		if j < n {
			out = Min(out, j-1-i)
		}
	}

	for j := r; j < n; j++ { // 当后缀为 [j, n-1] 时， 找出符合非递减的前缀 [0, i]
		i := sort.Search(l+1, func(i int) bool { return arr[i] > arr[j] }) - 1
		if i < l {
			out = Min(out, j-1-i)
		}
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
