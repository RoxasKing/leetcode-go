package main

import "sort"

/*
  Given a positive integer n, find the smallest integer which has exactly the same digits existing in the integer n and is greater in value than n. If no such positive integer exists, return -1.

  Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does not fit in 32-bit integer, return -1.

  Example 1:
    Input: n = 12
    Output: 21

  Example 2:
    Input: n = 21
    Output: -1

  Constraints:
    1 <= n <= 231 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/next-greater-element-iii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func nextGreaterElement(n int) int {
	arr := []int{}
	for i := n; i > 0; i /= 10 {
		arr = append(arr, i%10)
	}

	reverse(arr)

	m := len(arr)
	i := m - 1
	for i-1 >= 0 && arr[i-1] >= arr[i] {
		i--
	}

	reverse(arr[i:])

	if i == 0 {
		return -1
	}

	j := sort.Search(m-i, func(j int) bool { return arr[j+i] > arr[i-1] }) + i
	arr[i-1], arr[j] = arr[j], arr[i-1]

	out := 0
	for _, num := range arr {
		if (1<<31-1-num)/10 < out {
			return -1
		}
		out = out*10 + num
	}
	return out
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
