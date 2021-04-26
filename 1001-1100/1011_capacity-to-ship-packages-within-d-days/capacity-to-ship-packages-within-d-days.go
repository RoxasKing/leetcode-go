package main

/*
  A conveyor belt has packages that must be shipped from one port to another within D days.

  The ith package on the conveyor belt has a weight of weights[i]. Each day, we load the ship with packages on the conveyor belt (in the order given by weights). We may not load more weight than the maximum weight capacity of the ship.

  Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within D days.

  Example 1:
    Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
    Output: 15
    Explanation: A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
      1st day: 1, 2, 3, 4, 5
      2nd day: 6, 7
      3rd day: 8
      4th day: 9
      5th day: 10

    Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.

  Example 2:
    Input: weights = [3,2,2,4,1,4], D = 3
    Output: 6
    Explanation: A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
      1st day: 3, 2
      2nd day: 2, 4
      3rd day: 1, 4

  Example 3:
    Input: weights = [1,2,3,1,1], D = 4
    Output: 3
    Explanation:
      1st day: 1
      2nd day: 2
      3rd day: 3
      4th day: 1, 1

  Constraints:
    1. 1 <= D <= weights.length <= 5 * 10^4
    2. 1 <= weights[i] <= 500

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func shipWithinDays(weights []int, D int) int {
	l, r := 0, 0
	for _, w := range weights {
		l = Max(l, w)
		r += w
	}
	for l < r {
		m := (l + r) >> 1
		if needDays(weights, m) > D {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func needDays(weights []int, limit int) int {
	out, cur := 1, 0
	for _, w := range weights {
		if cur+w > limit {
			out++
			cur = w
		} else {
			cur += w
		}
	}
	return out
}
