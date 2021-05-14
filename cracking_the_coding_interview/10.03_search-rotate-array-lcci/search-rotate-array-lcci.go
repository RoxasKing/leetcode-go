package main

/*
  Given a sorted array of n integers that has been rotated an unknown number of times, write code to find an element in the array. You may assume that the array was originally sorted in increasing order. If there are more than one target elements in the array, return the smallest index.

  Example1:
     Input: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 5
     Output: 8 (the index of 5 in the array)

  Example2:
    Input: arr = [15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14], target = 11
    Output: -1 (not found)

  Note:
    1 <= arr.length <= 1000000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/search-rotate-array-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search

func search(arr []int, target int) int {
	n := len(arr)
	l, r := 0, n-1
	for l < r && arr[l] == arr[r] {
		r--
	}
	rotateIdx := bSearchRotateIdx(arr, l, r)
	if rotateIdx == 0 || arr[l] > target {
		return bSearch(arr, rotateIdx, r, target)
	}
	return bSearch(arr, l, rotateIdx, target)
}

func bSearchRotateIdx(arr []int, l, r int) int {
	start := l
	for l < r {
		m := (l + r) >> 1
		if arr[m] > arr[m+1] {
			return m + 1
		}
		if arr[l] <= arr[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return start
}

func bSearch(arr []int, l, r, target int) int {
	for l < r {
		m := (l + r) >> 1
		if arr[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if arr[l] == target {
		return l
	}
	return -1
}
