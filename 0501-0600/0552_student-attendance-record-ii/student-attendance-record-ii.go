package main

// Tags:
// Dynamic Programming

const mod int = 1e9 + 7

func checkRecord(n int) int {
	f0 := [2][3]int{}
	f0[0][0] = 1
	f0[1][0] = 1
	f0[0][1] = 1

	for n--; n > 0; n-- {
		f1 := [2][3]int{}
		for i := 0; i < 2; i++ { // Absent days
			for j := 0; j < 3; j++ { // consecutive Late days
				if i+1 < 2 { // A
					f1[i+1][0] += f0[i][j]
					f1[i+1][0] %= mod
				}
				if j+1 < 3 { // L
					f1[i][j+1] += f0[i][j]
					f1[i][j+1] %= mod
				}
				f1[i][0] += f0[i][j] // P
				f1[i][0] %= mod
			}
		}
		f0 = f1
	}

	var out int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			out = (out + f0[i][j]) % mod
		}
	}
	return out
}
