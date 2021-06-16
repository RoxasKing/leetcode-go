package main

/*
  Let's call an array arr a mountain if the following properties hold:

    1. arr.length >= 3
    2. There exists some i with 0 < i < arr.length - 1 such that:
       1. arr[0] < arr[1] < ... arr[i-1] < arr[i]
       2. arr[i] > arr[i+1] > ... > arr[arr.length - 1]

  Given an integer array arr that is guaranteed to be a mountain, return any i such that arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1].

  Example 1:
    Input: arr = [0,1,0]
    Output: 1

  Example 2:
    Input: arr = [0,2,1,0]
    Output: 1

  Example 3:
    Input: arr = [0,10,5,2]
    Output: 1

  Example 4:
    Input: arr = [3,4,5,1]
    Output: 2

  Example 5:
    Input: arr = [24,69,100,99,79,78,67,36,26,19]
    Output: 2

  Constraints:
    1. 3 <= arr.length <= 10^4
    2. 0 <= arr[i] <= 10^6
    3. arr is guaranteed to be a mountain array.

  Follow up: Finding the O(n) is straightforward, could you find an O(log(n)) solution?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/peak-index-in-a-mountain-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l+1 < r {
		m := (l + r) >> 1
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		}
		if arr[m] <= arr[m+1] {
			l = m
		} else {
			r = m
		}
	}
	return -1
}
