package My_LeetCode_In_Go

/*
  给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	l, r := 0, len(height)-1
	lmax, rmax := height[l], height[r]
	var out int
	for l <= r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])
		if lmax < rmax {
			out += lmax - height[l]
			l++
		} else {
			out += rmax - height[r]
			r--
		}
	}
	return out
}
