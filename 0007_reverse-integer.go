package My_LeetCode_In_Go

/*
  给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
*/

func reverse(x int) int {
	var out = 0
	for ; x != 0; x /= 10 {
		if out < (-1<<31)/10 || out > (1<<31-1)/10 {
			return 0
		}
		out = out*10 + x%10
	}
	return out
}
