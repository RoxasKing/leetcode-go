package interview

import "math"

/*
  给定两条线段（表示为起点start = {X1, Y1}和终点end = {X2, Y2}），如果它们有交点，请计算其交点，没有交点则返回空值。
  要求浮点型误差不超过10^-6。若有多个交点（线段重叠）则返回 X 值最小的点，X 坐标相同则返回 Y 值最小的点。
*/

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	var ans []float64

	x1, y1 := float64(start1[0]), float64(start1[1])
	x2, y2 := float64(end1[0]), float64(end1[1])
	x3, y3 := float64(start2[0]), float64(start2[1])
	x4, y4 := float64(end2[0]), float64(end2[1])

	inside := func(x1, y1, x2, y2, xk, yk float64) bool {
		return (x1 == x2 || math.Min(x1, x2) <= xk && xk <= math.Max(x1, x2)) &&
			(y1 == y2 || math.Min(y1, y2) <= yk && yk <= math.Max(y1, y2))
	}

	update := func(xk, yk float64) {
		if ans == nil || xk < ans[0] || xk == ans[0] && yk < ans[1] {
			ans = []float64{xk, yk}
		}
	}

	if (y4-y3)*(x2-x1) == (y2-y1)*(x4-x3) {
		if (y2-y1)*(x3-x1) == (y3-y1)*(x2-x1) {
			if inside(x1, y1, x2, y2, x3, y3) {
				update(x3, y3)
			}
			if inside(x1, y1, x2, y2, x4, y4) {
				update(x4, y4)
			}
			if inside(x3, y3, x4, y4, x1, y1) {
				update(x1, y1)
			}
			if inside(x3, y3, x4, y4, x2, y2) {
				update(x2, y2)
			}
		}
	} else {
		t1 := (x3*(y4-y3) + y1*(x4-x3) - y3*(x4-x3) - x1*(y4-y3)) /
			((x2-x1)*(y4-y3) - (x4-x3)*(y2-y1))
		t2 := (x1*(y2-y1) + y3*(x2-x1) - y1*(x2-x1) - x3*(y2-y1)) /
			((x4-x3)*(y2-y1) - (x2-x1)*(y4-y3))
		if t1 >= 0 && t1 <= 1 && t2 >= 0 && t2 <= 1 {
			ans = []float64{x1 + t1*(x2-x1), y1 + t1*(y2-y1)}
		}

	}
	return ans
}
