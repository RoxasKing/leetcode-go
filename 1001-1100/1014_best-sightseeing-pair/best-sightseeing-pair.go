package main

/*
  You are given an integer array values where values[i] represents the value of the ith sightseeing spot. Two sightseeing spots i and j have a distance j - i between them.

  The score of a pair (i < j) of sightseeing spots is values[i] + values[j] + i - j: the sum of the values of the sightseeing spots, minus the distance between them.

  Return the maximum score of a pair of sightseeing spots.

  Example 1:
    Input: values = [8,1,5,2,6]
    Output: 11
    Explanation: i = 0, j = 2, values[i] + values[j] + i - j = 8 + 5 + 0 - 2 = 11

  Example 2:
    Input: values = [1,2]
    Output: 2

  Constraints:
    1. 2 <= values.length <= 5 * 10^4
    2. 1 <= values[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/best-sightseeing-pair
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	out := 0
	max := 0
	for i := 0; i < n; i++ {
		out = Max(out, max+values[i]-i) // (values[i] + i) + (values[j] - j)
		max = Max(max, values[i]+i)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
