package My_LeetCode_In_Go

/*
  给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
  最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
  你可以假设除了整数 0 之外，这个整数不会以零开头。
*/

func plusOne(digits []int) []int {
	cur := 1
	out := make([]int, 0, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		out = append(out, (digits[i]+cur)%10)
		cur = (digits[i] + cur) / 10
	}
	if cur > 0 {
		out = append(out, cur)
	}
	for i := 0; i < len(out)/2; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}
