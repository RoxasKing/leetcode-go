package main

/*
  There is a biker going on a road trip. The road trip consists of n + 1 points at different altitudes. The biker starts his trip on point 0 with altitude equal 0.

  You are given an integer array gain of length n where gain[i] is the net gain in altitude between points i​​​​​​ and i + 1 for all (0 <= i < n). Return the highest altitude of a point.

  Example 1:
    Input: gain = [-5,1,5,0,-7]
    Output: 1
    Explanation: The altitudes are [0,-5,-4,1,1,-6]. The highest is 1.

  Example 2:
    Input: gain = [-4,-3,-2,-1,4,3,2]
    Output: 0
    Explanation: The altitudes are [0,-4,-7,-9,-10,-6,-3,-1]. The highest is 0.

  Constraints:
    1. n == gain.length
    2. 1 <= n <= 100
    3. -100 <= gain[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-the-highest-altitude
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func largestAltitude(gain []int) int {
	out, sum := 0, 0
	for _, g := range gain {
		sum += g
		out = Max(out, sum)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
