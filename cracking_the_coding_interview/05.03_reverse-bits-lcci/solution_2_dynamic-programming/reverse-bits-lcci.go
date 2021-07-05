package main

// Tags:
// Dynamic Programming

func reverseBits(num int) int {
	out := 0
	x := uint(num)
	f := [32][2]int{}
	f[0] = [2]int{int(x & 1), 1}
	for i := 1; i <= 31; i++ {
		if (x>>i)&1 == 1 {
			f[i] = [2]int{f[i-1][0] + 1, f[i-1][1] + 1}
		} else {
			f[i][1] = f[i-1][0] + 1
		}
		out = Max(out, Max(f[i][0], f[i][1]))
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
