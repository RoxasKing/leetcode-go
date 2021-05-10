package main

/*
  You are given a sorted integer array arr containing 1 and prime numbers, where all the integers of arr are unique. You are also given an integer k.

  For every i and j where 0 <= i < j < arr.length, we consider the fraction arr[i] / arr[j].

  Return the kth smallest fraction considered. Return your answer as an array of integers of size 2, where answer[0] == arr[i] and answer[1] == arr[j].

  Example 1:
    Input: arr = [1,2,3,5], k = 3
    Output: [2,5]
    Explanation: The fractions to be considered in sorted order are:
      1/5, 1/3, 2/5, 1/2, 3/5, and 2/3.
      The third fraction is 2/5.

  Example 2:
    Input: arr = [1,7], k = 1
    Output: [1,7]

  Constraints:
    1. 2 <= arr.length <= 1000
    2. 1 <= arr[i] <= 3 * 10^4
    3. arr[0] == 1
    4. arr[i] is a prime number for i > 0.
    5. All the numbers of arr are unique and sorted in strictly increasing order.
    6. 1 <= k <= arr.length * (arr.length - 1) / 2

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/k-th-smallest-prime-fraction
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Binary Search
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	a, b := 0, 1
	l, r := float64(0), float64(1)
	for r-l > 1e-9 {
		target := (l + r) / 2.0
		cnt, x, y := countSmaller(arr, n, target)
		if cnt < k {
			l = target
		} else {
			r = target
			a, b = x, y
		}
	}
	return []int{a, b}
}

func countSmaller(arr []int, n int, target float64) (int, int, int) {
	cnt, a, b := 0, 0, 1
	for i, j := -1, 1; j < n; j++ {
		for float64(arr[i+1])/float64(arr[j]) < target {
			i++
		}
		cnt += i + 1
		if i >= 0 && float64(a)/float64(b) < float64(arr[i])/float64(arr[j]) {
			a, b = arr[i], arr[j]
		}
	}
	return cnt, a, b
}
