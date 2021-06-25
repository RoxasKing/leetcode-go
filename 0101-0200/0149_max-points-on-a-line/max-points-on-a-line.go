package main

/*
  Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane, return the maximum number of points that lie on the same straight line.

  Example 1:
    Input: points = [[1,1],[2,2],[3,3]]
    Output: 3

  Example 2:
    Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
    Output: 4

  Constraints:
    1. 1 <= points.length <= 300
    2. points[i].length == 2
    3. -10^4 <= xi, yi <= 10^4
    4. All the points are unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/max-points-on-a-line
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash

func maxPoints(points [][]int) int {
	out := 1
	cntX := map[int]int{}
	cnty := map[int]int{}
	for i, p1 := range points {
		cntX[p1[0]]++
		cnty[p1[1]]++
		for j, p2 := range points[i+1:] {
			if p2[0] == p1[0] || p2[1] == p1[1] {
				continue
			}
			cnt := 2
			for _, p3 := range points[j+1:] {
				if p3[0] == p1[0] || p3[1] == p1[1] || p3[0] == p2[0] || p3[1] == p2[1] {
					continue
				}
				if (p2[0]-p1[0])*(p3[1]-p1[1]) == (p3[0]-p1[0])*(p2[1]-p1[1]) {
					cnt++
				}
			}
			out = Max(out, cnt)
		}
	}
	for _, cnt := range cntX {
		out = Max(out, cnt)
	}
	for _, cnt := range cnty {
		out = Max(out, cnt)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
