package My_LeetCode_In_Go

/*
  实现 int sqrt(int x) 函数。
  计算并返回 x 的平方根，其中 x 是非负整数。
  由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
*/

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)>>1
		switch {
		case mid*mid < x:
			l = mid + 1
		case mid*mid > x:
			r = mid - 1
		case mid*mid == x:
			return mid
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
