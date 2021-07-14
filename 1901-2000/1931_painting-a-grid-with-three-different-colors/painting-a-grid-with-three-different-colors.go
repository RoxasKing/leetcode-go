package main

// Tags:
// Dynamic Programming

func colorTheGrid(m int, n int) int {
	base := make([]int, m)
	base[0] = 1
	for i := 1; i < m; i++ {
		base[i] = base[i-1] * 3
	}

	arr := []int{}
	for i := 0; i < base[m-1]*3; i++ {
		ok := true
		for j := 1; j < m; j++ {
			if i/base[j]%3 == i/base[j-1]%3 {
				ok = false
				break
			}
		}
		if ok {
			arr = append(arr, i)
		}
	}

	k := len(arr)
	validPairs := make([][]int, k)
	for i := 0; i < k; i++ {
		x := arr[i]
		for j := i + 1; j < k; j++ {
			y := arr[j]
			ok := true
			for k := 0; k < m; k++ {
				if x/base[k]%3 == y/base[k]%3 {
					ok = false
					break
				}
			}
			if ok {
				validPairs[i] = append(validPairs[i], j)
				validPairs[j] = append(validPairs[j], i)
			}
		}
	}

	f0 := make([]int, k)
	f1 := make([]int, k)
	for i := range f0 {
		f0[i] = 1
	}

	for ; n > 1; n-- {
		for i := 0; i < k; i++ {
			f1[i] = 0
			for _, j := range validPairs[i] {
				f1[i] = (f1[i] + f0[j]) % mod
			}
		}
		f0, f1 = f1, f0
	}

	var out int
	for i := 0; i < k; i++ {
		out = (out + f0[i]) % mod
	}
	return out
}

var mod = int(1e9 + 7)
