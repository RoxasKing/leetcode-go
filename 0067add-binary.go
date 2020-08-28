package main

/*
  给定两个二进制字符串，返回他们的和（用二进制表示）。
  输入为非空字符串且只包含数字 1 和 0。
*/

func addBinary(a string, b string) string {
	out := make([]byte, 0, len(a)+len(b)+1)
	var cur byte
	for a != "" && b != "" {
		sum := a[len(a)-1] - '0' + b[len(b)-1] - '0' + cur
		out = append(out, sum%2+'0')
		cur = sum >> 1
		a = a[:len(a)-1]
		b = b[:len(b)-1]
	}
	for a != "" {
		sum := a[len(a)-1] - '0' + cur
		out = append(out, sum%2+'0')
		cur = sum >> 1
		a = a[:len(a)-1]
	}
	for b != "" {
		sum := b[len(b)-1] - '0' + cur
		out = append(out, sum%2+'0')
		cur = sum >> 1
		b = b[:len(b)-1]
	}
	if cur != 0 {
		out = append(out, cur+'0')
	}
	for i := 0; i < len(out)/2; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return string(out)
}
