package main

import "sort"

/*
  Given an integer array arr and a target value target, return the integer value such that when we change all the integers larger than value in the given array to be equal to value, the sum of the array gets as close as possible (in absolute difference) to target.

  In case of a tie, return the minimum such integer.

  Notice that the answer is not neccesarilly a number from arr.

  Example 1:
    Input: arr = [4,9,3], target = 10
    Output: 3
    Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's the optimal answer.

  Example 2:
    Input: arr = [2,3,5], target = 10
    Output: 5

  Example 3:
    Input: arr = [60864,25176,27249,21296,20204], target = 56803
    Output: 11361

  Constraints:
    1. 1 <= arr.length <= 10^4
    2. 1 <= arr[i], target <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Binary Search
// Prefix Sum

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + arr[i]
	}

	l, r := 0, arr[n-1]
	for l < r {
		m := l + (r-l)>>1
		i := sort.SearchInts(arr, m)
		sum := pSum[i] + m*(n-i)
		if sum < target {
			l = m + 1
		} else {
			r = m
		}
	}

	i := sort.SearchInts(arr, l)
	j := sort.SearchInts(arr, l-1)
	if Abs((n-i)*l+pSum[i]-target) < Abs((n-j)*(l-1)+pSum[j]-target) {
		return l
	}
	return l - 1
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
