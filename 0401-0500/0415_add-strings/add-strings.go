package main

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
