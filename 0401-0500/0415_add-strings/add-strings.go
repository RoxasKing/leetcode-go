package main

/*
  Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

  Example 1:
    Input: num1 = "11", num2 = "123"
    Output: "134"

  Example 2:
    Input: num1 = "456", num2 = "77"
    Output: "533"

  Example 3:
    Input: num1 = "0", num2 = "0"
    Output: "0"

  Constraints:
    1. 1 <= num1.length, num2.length <= 10^4
    2. num1 and num2 consist of only digits.
    3. num1 and num2 don't have any leading zeros except for the zero itself.

  Follow up: Could you solve it without using any built-in BigInteger library or converting the inputs to integer directly?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/add-strings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	chs := make([]byte, 0, Max(m, n)+1)
	remain := 0
	for i, j := m-1, n-1; i >= 0 || j >= 0; {
		if i >= 0 {
			remain += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			remain += int(num2[j] - '0')
			j--
		}
		chs = append(chs, byte(remain%10)+'0')
		remain /= 10
	}
	if remain > 0 {
		chs = append(chs, byte(remain)+'0')
	}
	for i := 0; i < len(chs)>>1; i++ {
		chs[i], chs[len(chs)-1-i] = chs[len(chs)-1-i], chs[i]
	}
	return string(chs)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
