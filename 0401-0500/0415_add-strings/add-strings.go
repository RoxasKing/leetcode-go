package main

func addStrings(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	nums := make([]byte, 0, Max(m, n)+1)
	remain := 0
	for i := 0; i < Max(m, n); i++ {
		if i < m {
			remain += int(num1[m-1-i] - '0')
		}
		if i < n {
			remain += int(num2[n-1-i] - '0')
		}
		nums = append(nums, byte(remain%10)+'0')
		remain /= 10
	}
	if remain > 0 {
		nums = append(nums, byte(remain%10)+'0')
	}
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	return string(nums)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
