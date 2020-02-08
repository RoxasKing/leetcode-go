package leetcode

/*
  给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
  说明：
    1.num1 和 num2 的长度小于110。
    2.num1 和 num2 只包含数字 0-9。
    3.num1 和 num2 均不以零开头，除非是数字 0 本身。
    4.不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	result := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			sum := result[i+j+1] + int(num1[i]-'0')*int(num2[j]-'0')
			result[i+j+1] = sum % 10
			result[i+j] += sum / 10
		}
	}
	out := make([]rune, 0, len(result))
	for i := range result {
		if i == 0 && result[i] == 0 {
			continue
		}
		out = append(out, rune(result[i]+'0'))
	}
	return string(out)
}
