package main

// Tags:
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
