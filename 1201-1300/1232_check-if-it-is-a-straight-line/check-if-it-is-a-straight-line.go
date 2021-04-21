package main

/*
  You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.

  Example 1:
    Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
    Output: true

  Example 2:
    Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
    Output: false

  Constraints:
    2 <= coordinates.length <= 1000
    coordinates[i].length == 2
    -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
    coordinates contains no duplicate point.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func checkStraightLine(coordinates [][]int) bool {
	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]
	for i := 2; i < len(coordinates); i++ {
		xi, yi := coordinates[i][0], coordinates[i][1]
		if (y1-y0)*(xi-x0) != (yi-y0)*(x1-x0) {
			return false
		}
	}
	return true
}
