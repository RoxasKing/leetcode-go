package My_LeetCode_In_Go

/*
  判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var flip int
	for i := x; i > 0; i /= 10 {
		flip = flip*10 + i%10
	}
	if flip != x {
		return false
	}
	return true
}
