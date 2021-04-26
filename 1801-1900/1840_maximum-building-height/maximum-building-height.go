package main

import (
	"sort"
)

/*
  You want to build n new buildings in a city. The new buildings will be built in a line and are labeled from 1 to n.

  However, there are city restrictions on the heights of the new buildings:
    1. The height of each building must be a non-negative integer.
    2. The height of the first building must be 0.
    3. The height difference between any two adjacent buildings cannot exceed 1.

  Additionally, there are city restrictions on the maximum height of specific buildings. These restrictions are given as a 2D integer array restrictions where restrictions[i] = [idi, maxHeighti] indicates that building idi must have a height less than or equal to maxHeighti.

  It is guaranteed that each building will appear at most once in restrictions, and building 1 will not be in restrictions.

  Return the maximum possible height of the tallest building.



  Example 1:
    Input: n = 5, restrictions = [[2,1],[4,1]]
    Output: 2
    Explanation: The green area in the image indicates the maximum allowed height for each building.
      We can build the buildings with heights [0,1,2,1,2], and the tallest building has a height of 2.

  Example 2:
    Input: n = 6, restrictions = []
    Output: 5
    Explanation: The green area in the image indicates the maximum allowed height for each building.
      We can build the buildings with heights [0,1,2,3,4,5], and the tallest building has a height of 5.

  Example 3:
    Input: n = 10, restrictions = [[5,3],[2,5],[7,4],[10,3]]
    Output: 5
    Explanation: The green area in the image indicates the maximum allowed height for each building.
      We can build the buildings with heights [0,1,2,3,3,4,4,5,4,3], and the tallest building has a height of 5.

  Constraints:
    1. 2 <= n <= 10^9
    2. 0 <= restrictions.length <= min(n - 1, 10^5)
    3. 2 <= id^i <= n
    4. id^i is unique.
    5. 0 <= maxHeight^i <= 109

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-building-height
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algorithm
func maxBuilding(n int, restrictions [][]int) int {
	m := len(restrictions)
	if m == 0 {
		return n - 1
	}

	sort.Slice(restrictions, func(i, j int) bool { return restrictions[i][0] < restrictions[j][0] })
	k := m + 1
	if restrictions[m-1][0] < n {
		k++
	}
	arr := make([][]int, k)
	copy(arr[1:1+m], restrictions)
	arr[0] = []int{1, 0}
	if k == m+2 {
		arr[k-1] = []int{n, n - 1}
	}

	for i := 1; i < k; i++ {
		h := Min(arr[i][1], arr[i-1][1]) + (arr[i][0] - arr[i-1][0])
		if h <= Max(arr[i][1], arr[i-1][1]) {
			arr[i][1] = Min(arr[i][1], h)
		}
	}
	out := arr[k-1][1]
	for i := k - 2; i >= 1; i-- {
		h := Min(arr[i][1], arr[i+1][1]) + (arr[i+1][0] - arr[i][0])
		if h <= Max(arr[i][1], arr[i+1][1]) {
			arr[i][1] = Min(arr[i][1], h)
			out = Max(out, arr[i][1])
		} else {
			out = Max(out, Max(arr[i][1], arr[i+1][1])+(arr[i+1][0]-(arr[i][0]+Abs(arr[i+1][1]-arr[i][1])))>>1)
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
