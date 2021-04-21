package main

import (
	"math"
	"sort"
)

/*
  You are given an array points, an integer angle, and your location, where location = [posx, posy] and points[i] = [xi, yi] both denote integral coordinates on the X-Y plane.

  Initially, you are facing directly east from your position. You cannot move from your position, but you can rotate. In other words, posx and posy cannot be changed. Your field of view in degrees is represented by angle, determining how wide you can see from any given view direction. Let d be the amount in degrees that you rotate counterclockwise. Then, your field of view is the inclusive range of angles [d - angle/2, d + angle/2].

  You can see some set of points if, for each point, the angle formed by the point, your position, and the immediate east direction from your position is in your field of view.

  There can be multiple points at one coordinate. There may be points at your location, and you can always see these points regardless of your rotation. Points do not obstruct your vision to other points.

  Return the maximum number of points you can see.

  Example 1:
    Input: points = [[2,1],[2,2],[3,3]], angle = 90, location = [1,1]
    Output: 3
    Explanation: The shaded region represents your field of view. All points can be made visible in your field of view, including [3,3] even though [2,2] is in front and in the same line of sight.

  Example 2:
    Input: points = [[2,1],[2,2],[3,4],[1,1]], angle = 90, location = [1,1]
    Output: 4
    Explanation: All points can be made visible in your field of view, including the one at your location.

  Example 3:
    Input: points = [[1,0],[2,1]], angle = 13, location = [1,1]
    Output: 1
    Explanation: You can only see one of the two points, as shown above.

  Constraints:
    1. 1 <= points.length <= 10^5
    2. points[i].length == 2
    3. location.length == 2
    4. 0 <= angle < 360
    5. 0 <= posx, posy, xi, yi <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-number-of-visible-points
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointerss
func visiblePoints(points [][]int, angle int, location []int) int {
	duplicate := 0
	angles := []float64{}
	for _, p := range points {
		x, y := p[0]-location[0], p[1]-location[1]
		if x == 0 && y == 0 {
			duplicate++
			continue
		}
		angle := math.Atan2(float64(y), float64(x)) * 180 / math.Pi
		angles = append(angles, angle)
	}
	sort.Float64s(angles)
	for _, angle := range angles {
		angles = append(angles, angle+360)
	}
	max := 0
	for l, r := 0, 0; l < len(angles); l++ {
		for r < len(angles) && (angles[r]-angles[l])-float64(angle) < 1e-8 {
			r++
		}
		if r-l > max {
			max = r - l
		}
	}
	return max + duplicate
}
