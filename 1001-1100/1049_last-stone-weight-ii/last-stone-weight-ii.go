package main

// Tags:
// Dynamic Programming

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	m := sum >> 1
	f0 := make([]bool, m+1)
	f1 := make([]bool, m+1)
	f0[0] = true

	for _, stone := range stones {
		for j := 0; j <= m; j++ {
			f1[j] = f0[j]
			if j >= stone {
				f1[j] = f1[j] || f0[j-stone]
			}
		}
		f0, f1 = f1, f0
	}

	for j := m; ; j-- {
		if f0[j] {
			return sum - 2*j
		}
	}
}
