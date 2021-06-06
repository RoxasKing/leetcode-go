package main

/*
  Given the coordinates of two rectilinear rectangles in a 2D plane, return the total area covered by the two rectangles.

  The first rectangle is defined by its bottom-left corner (ax1, ay1) and its top-right corner (ax2, ay2).

  The second rectangle is defined by its bottom-left corner (bx1, by1) and its top-right corner (bx2, by2).

  Example 1:
    Input: ax1 = -3, ay1 = 0, ax2 = 3, ay2 = 4, bx1 = 0, by1 = -1, bx2 = 9, by2 = 2
    Output: 45

  Example 2:
    Input: ax1 = -2, ay1 = -2, ax2 = 2, ay2 = 2, bx1 = -2, by1 = -2, bx2 = 2, by2 = 2
    Output: 16

  Constraints:
    -10^4 <= ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/rectangle-area
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	out := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	l, r := Max(ax1, bx1), Min(ax2, bx2)
	u, d := Min(ay2, by2), Max(ay1, by1)
	if l < r && u > d {
		out -= (r - l) * (u - d)
	}
	return out
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
