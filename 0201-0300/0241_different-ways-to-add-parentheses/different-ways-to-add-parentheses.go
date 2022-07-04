package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func diffWaysToCompute(expression string) []int {
	arr := []int{}
	ops := []rune{}
	num := 0
	for _, c := range expression {
		if '0' <= c && c <= '9' {
			num = num*10 + int(c-'0')
			continue
		}
		arr = append(arr, num)
		num = 0
		ops = append(ops, c)
	}
	arr = append(arr, num)
	n := len(arr)
	f := make([][][]int, n)
	for i := range f {
		f[i] = make([][]int, n)
		f[i][i] = append(f[i][i], arr[i])
	}
	for k := 1; k < n; k++ {
		for l, r := 0, k; r < n; l, r = l+1, r+1 {
			for i := l; i < r; i++ {
				lvs, rvs := f[l][i], f[i+1][r]
				for _, lv := range lvs {
					for _, rv := range rvs {
						v := 0
						if ops[i] == '+' {
							v = lv + rv
						} else if ops[i] == '-' {
							v = lv - rv
						} else {
							v = lv * rv
						}
						f[l][r] = append(f[l][r], v)
					}
				}
			}
		}
	}
	return f[0][n-1]
}
