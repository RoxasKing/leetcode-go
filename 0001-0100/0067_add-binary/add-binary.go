package main

/*
  给你两个二进制字符串，返回它们的和（用二进制表示）。
  输入为 非空 字符串且只包含数字 1 和 0。

  示例 1:
    输入: a = "11", b = "1"
    输出: "100"

  示例 2:
    输入: a = "1010", b = "1011"
    输出: "10101"

  提示：
    每个字符串仅由字符 '0' 或 '1' 组成。
    1 <= a.length, b.length <= 10^4
    字符串如果不是 "0" ，就都不含前导零。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/add-binary
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func addBinary(a string, b string) string {
	arrA, arrB := []byte(a), []byte(b)
	reverse(arrA)
	reverse(arrB)
	arr := []byte{}
	remain := 0
	for len(arrA) > 0 && len(arrB) > 0 {
		remain += int(arrA[0]-'0') + int(arrB[0]-'0')
		arr = append(arr, byte(remain%2+'0'))
		remain /= 2
		arrA = arrA[1:]
		arrB = arrB[1:]
	}
	for len(arrA) > 0 {
		remain += int(arrA[0] - '0')
		arr = append(arr, byte(remain%2+'0'))
		remain /= 2
		arrA = arrA[1:]
	}
	for len(arrB) > 0 {
		remain += int(arrB[0] - '0')
		arr = append(arr, byte(remain%2+'0'))
		remain /= 2
		arrB = arrB[1:]
	}
	if remain > 0 {
		arr = append(arr, byte(remain+'0'))
	}
	reverse(arr)
	return string(arr)
}

func reverse(arr []byte) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
