package main

/*
  给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
*/

func reverse(x int) int {
	var out int
	for x != 0 {
		if out < (-1<<31)/10 || (1<<31-1)/10 < out {
			return 0
		}
		out = out*10 + x%10
		x /= 10
	}
	return out
}
