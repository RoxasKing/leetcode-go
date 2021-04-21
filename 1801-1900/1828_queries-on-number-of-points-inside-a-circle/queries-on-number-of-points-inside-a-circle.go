package main

/*
  You are given an array points where points[i] = [xi, yi] is the coordinates of the ith point on a 2D plane. Multiple points can have the same coordinates.

  You are also given an array queries where queries[j] = [xj, yj, rj] describes a circle centered at (xj, yj) with a radius of rj.

  For each query queries[j], compute the number of points inside the jth circle. Points on the border of the circle are considered inside.

  Return an array answer, where answer[j] is the answer to the jth query.

  Example 1:
    Input: points = [[1,3],[3,3],[5,3],[2,2]], queries = [[2,3,1],[4,3,1],[1,1,2]]
    Output: [3,2,2]
    Explanation: The points and circles are shown above.
    queries[0] is the green circle, queries[1] is the red circle, and queries[2] is the blue circle.

  Example 2:
    Input: points = [[1,1],[2,2],[3,3],[4,4],[5,5]], queries = [[1,2,2],[2,2,2],[4,3,2],[4,3,3]]
    Output: [2,3,2,4]
    Explanation: The points and circles are shown above.
    queries[0] is green, queries[1] is red, queries[2] is blue, and queries[3] is purple.

  Constraints:
    1. 1 <= points.length <= 500
    2. points[i].length == 2
    3. 0 <= x​​​​​​i, y​​​​​​i <= 500
    4. 1 <= queries.length <= 500
    5. queries[j].length == 3
    6. 0 <= xj, yj <= 500
    7. 1 <= rj <= 500
    8. All coordinates are integers.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/queries-on-number-of-points-inside-a-circle
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countPoints(points [][]int, queries [][]int) []int {
	n := len(queries)
	out := make([]int, n)
	for i, q := range queries {
		for _, p := range points {
			if getDist(q[0], q[1], p[0], p[1]) <= q[2]*q[2] {
				out[i]++
			}
		}
	}
	return out
}

func getDist(x1, y1, x2, y2 int) int {
	return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
}
