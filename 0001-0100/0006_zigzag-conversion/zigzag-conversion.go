package main

// Difficulty:
// Medium

func convert(s string, numRows int) string {
	n := len(s)
	if n <= numRows || 1 == numRows {
		return s
	}
	arr := make([]byte, 0, n)
	for row := 0; row < numRows; row++ {
		for i := row; i < n; i += (numRows - 1) << 1 {
			arr = append(arr, s[i])
			if row == 0 || row == numRows-1 {
				continue
			}
			if j := i + (numRows-1-row)<<1; j < len(s) {
				arr = append(arr, s[j])
			}
		}
	}
	return string(arr)
}
