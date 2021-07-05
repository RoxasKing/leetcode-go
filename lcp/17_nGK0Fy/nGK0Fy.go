package main

func calculate(s string) int {
	x, y := 1, 0
	for i := range s {
		if s[i] == 'A' {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}
	return x + y
}
