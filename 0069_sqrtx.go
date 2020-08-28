package main

/*
  实现 int sqrt(int x) 函数。
  计算并返回 x 的平方根，其中 x 是非负整数。
  由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/

// Binary Search
func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		m := l + (r-l)>>1
		tmp := m * m
		if tmp < x {
			l = m + 1
		} else if tmp > x {
			r = m - 1
		} else {
			return m
		}
	}
	return l - 1
}

// Newton Method
func mySqrt2(x int) int {
	out := x
	for out*out > x {
		out = (out + x/out) / 2
	}
	return out
}
