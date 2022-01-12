package main

// Difficulty:
// Easy

// Tags:
// Simulation
// Two Pointers
// Bit Manipulation

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	remain := 0
	arr := make([]byte, 0, Max(m, n)+1)
	for i, j := m-1, n-1; i >= 0 || j >= 0; remain >>= 1 {
		if i >= 0 {
			remain += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			remain += int(b[j] - '0')
			j--
		}
		arr = append(arr, byte(remain&1)+'0')
	}
	if remain != 0 {
		arr = append(arr, byte(remain)+'0')
	}
	for i := 0; i < len(arr)>>1; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
