package main

/*
  Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

  Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

  Example 1:
    Input: num1 = "2", num2 = "3"
    Output: "6"

  Example 2:
    Input: num1 = "123", num2 = "456"
    Output: "56088"

  Constraints:
    1. 1 <= num1.length, num2.length <= 200
    2. num1 and num2 consist of digits only.
    3. Both num1 and num2 do not contain any leading zero, except the number 0 itself.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/multiply-strings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math
func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}

	m, n := len(num1), len(num2)
	arr := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			res := arr[i+j+1] + int(num1[i]-'0')*int(num2[j]-'0')
			arr[i+j+1] = res % 10
			arr[i+j] += res / 10
		}
	}

	if arr[0] == 0 {
		arr = arr[1:]
	}

	out := ""
	for _, num := range arr {
		out += string(byte(num) + '0')
	}
	return out
}
