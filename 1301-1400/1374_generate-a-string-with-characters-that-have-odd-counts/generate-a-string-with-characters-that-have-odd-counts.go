package main

// Difficulty:
// Easy

func generateTheString(n int) string {
	o := make([]byte, 0, n)
	for ch := byte('a'); n > 0; ch++ {
		t := 25
		for ; t > n; t -= 2 {
		}
		n -= t
		for ; t > 0; t-- {
			o = append(o, ch)
		}
	}
	return string(o)
}
