package main

func convert(s string, numRows int) string {
	n := len(s)
	if numRows >= n || numRows == 1 {
		return s
	}
	out := make([]byte, 0, n)
	for i := 0; i < numRows; i++ {
		for j := i; j < n; j += (numRows - 1) * 2 {
			out = append(out, s[j])
			if i == 0 || i == numRows-1 {
				continue
			}
			next := j + (numRows-1-i)*2
			if next < len(s) {
				out = append(out, s[next])
			}
		}
	}
	return string(out)
}
