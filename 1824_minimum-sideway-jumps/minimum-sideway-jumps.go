package main

/*
  There is a 3 lane road of length n that consists of n + 1 points labeled from 0 to n. A frog starts at point 0 in the second lane and wants to jump to point n. However, there could be obstacles along the way.

  You are given an array obstacles of length n + 1 where each obstacles[i] (ranging from 0 to 3) describes an obstacle on the lane obstacles[i] at point i. If obstacles[i] == 0, there are no obstacles at point i. There will be at most one obstacle in the 3 lanes at each point.
    For example, if obstacles[2] == 1, then there is an obstacle on lane 1 at point 2.
  The frog can only travel from point i to point i + 1 on the same lane if there is not an obstacle on the lane at point i + 1. To avoid obstacles, the frog can also perform a side jump to jump to another lane (even if they are not adjacent) at the same point if there is no obstacle on the new lane.
    For example, the frog can jump from lane 3 at point 3 to lane 1 at point 3.
  Return the minimum number of side jumps the frog needs to reach any lane at point n starting from lane 2 at point 0.

  Note: There will be no obstacles on points 0 and n.

  Example 1:
    Input: obstacles = [0,1,2,3,0]
    Output: 2
    Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps (red arrows).
    Note that the frog can jump over obstacles only when making side jumps (as shown at point 2).

  Example 2:
    Input: obstacles = [0,1,1,3,3,0]
    Output: 0
    Explanation: There are no obstacles on lane 2. No side jumps are required.

  Example 3:
    Input: obstacles = [0,2,1,0,3,0]
    Output: 2
    Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps.

  Constraints:
    1. obstacles.length == n + 1
    2. 1 <= n <= 5 * 10^5
    3. 0 <= obstacles[i] <= 3
    4. obstacles[0] == obstacles[n] == 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-sideway-jumps
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func minSideJumps(obstacles []int) int {
	dp0, dp1, dp2 := 1, 0, 1

	for i := 1; i < len(obstacles); i++ {
		dp00, dp01, dp02 := dp0, dp1, dp2
		switch obstacles[i] {
		case 1:
			dp0 = 1<<31 - 1
			dp1 = Min(dp01, dp02+1)
			dp2 = Min(dp02, dp01+1)
		case 2:
			dp0 = Min(dp00, dp02+1)
			dp1 = 1<<31 - 1
			dp2 = Min(dp02, dp00+1)
		case 3:
			dp0 = Min(dp00, dp01+1)
			dp1 = Min(dp01, dp00+1)
			dp2 = 1<<31 - 1
		default:
			dp0 = Min(dp00, Min(dp01+1, dp02+1))
			dp1 = Min(dp01, Min(dp00+1, dp02+1))
			dp2 = Min(dp02, Min(dp00+1, dp01+1))
		}
	}

	return Min(dp0, Min(dp1, dp2))
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
