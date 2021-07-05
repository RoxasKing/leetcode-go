package main

func minOperations(s string) int {
	x := 0
	min := 0
	for i := range s {
		if int(s[i]-'0') != x {
			min++
		}
		x ^= 1
	}

	x = 1
	cur := 0
	for i := range s {
		if int(s[i]-'0') != x {
			cur++
		}
		x ^= 1
	}

	min = Min(cur, min)

	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
