package util

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(n int) int {
	if n <= 0 {
		return 0 - n
	}
	return n
}
