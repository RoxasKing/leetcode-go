package main

// Difficulty:
// Medium

// Tags:
// Math
// Simulation

func multiply(num1 string, num2 string) string {
	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	m, n := len(num1), len(num2)
	nums := make([]byte, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			res := nums[i+j+1] + (num1[i]-'0')*(num2[j]-'0')
			nums[i+j+1] = res % 10
			nums[i+j] += res / 10
		}
	}
	if nums[0] == 0 {
		nums = nums[1:]
	}
	for i := range nums {
		nums[i] += '0'
	}
	return string(nums)
}
