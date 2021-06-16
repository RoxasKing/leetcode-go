package main

/*
  You are given an array of integers distance.

  You start at point (0,0) on an X-Y plane and you move distance[0] meters to the north, then distance[1] meters to the west, distance[2] meters to the south, distance[3] meters to the east, and so on. In other words, after each move, your direction changes counter-clockwise.

  Return true if your path crosses itself, and false if it does not.

  Example 1:
    Input: distance = [2,1,1,2]
    Output: true

  Example 2:
    Input: distance = [1,2,3,4]
    Output: false

  Example 3:
    Input: distance = [1,1,1,1]
    Output: true

  Constraints:
    1 <= distance.length <= 10^5
    1 <= distance[i] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/self-crossing
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math

func isSelfCrossing(dist []int) bool {
	if len(dist) < 4 {
		return false
	}
	i, n := 2, len(dist)
	for ; i < n && dist[i] > dist[i-2]; i++ {
	}
	if i == n {
		return false
	}

	if i == 3 && dist[i] == dist[i-2] || i >= 4 && dist[i] >= dist[i-2]-dist[i-4] {
		dist[i-1] -= dist[i-3]
	}

	for i++; i < n && dist[i] < dist[i-2]; i++ {
	}

	return i < n
}
