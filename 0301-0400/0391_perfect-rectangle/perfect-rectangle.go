package main

/*
  Given an array rectangles where rectangles[i] = [xi, yi, ai, bi] represents an axis-aligned rectangle. The bottom-left point of the rectangle is (xi, yi) and the top-right point of it is (ai, bi).

  Return true if all the rectangles together form an exact cover of a rectangular region.

  Example 1:
    Input: rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
    Output: true
    Explanation: All 5 rectangles together form an exact cover of a rectangular region.

  Example 2:
    Input: rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
    Output: false
    Explanation: Because there is a gap between the two rectangular regions.

  Example 3:
    Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[3,2,4,4]]
    Output: false
    Explanation: Because there is a gap in the top center.

  Example 4:
    Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
    Output: false
    Explanation: Because two of the rectangles overlap with each other.

  Constraints:
    1. 1 <= rectangles.length <= 2 * 10^4
    2. rectangles[i].length == 4
    3. -10^5 <= xi, yi, ai, bi <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/perfect-rectangle
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math
// Hash

func isRectangleCover(rectangles [][]int) bool {
	sumA := 0
	minX, minY, maxA, maxB := 1<<31-1, 1<<31-1, -1<<31, -1<<31
	freq := map[[2]int]int{}

	for _, rec := range rectangles {
		x, y, a, b := rec[0], rec[1], rec[2], rec[3]
		freq[[2]int{x, y}]++
		freq[[2]int{x, b}]++
		freq[[2]int{a, y}]++
		freq[[2]int{a, b}]++
		minX = Min(minX, x)
		minY = Min(minY, y)
		maxA = Max(maxA, a)
		maxB = Max(maxB, b)
		sumA += (a - x) * (b - y)
	}

	sumB := (maxA - minX) * (maxB - minY)

	if sumA != sumB {
		return false
	}

	for p, c := range freq {
		x, y := p[0], p[1]
		isCorner := (x == minX || x == maxA) && (y == minY || y == maxB)
		if isCorner && c > 1 || !isCorner && c&1 == 1 {
			return false
		}
	}

	return true
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
