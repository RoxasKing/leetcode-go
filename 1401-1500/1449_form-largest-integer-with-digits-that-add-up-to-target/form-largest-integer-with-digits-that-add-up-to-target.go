package main

// Tags:
// Dynamic Programming
// Knapsack Problem

func largestNumber(cost []int, target int) string {
	f := make([]int, target+1)
	for _, c := range cost {
		for i := c; i <= target; i++ {
			if i-c > 0 && f[i-c] == 0 {
				continue
			}
			f[i] = Max(f[i], f[i-c]+1)
		}
	}

	if f[target] == 0 {
		return "0"
	}

	out := make([]byte, 0, f[target])
	for i, j := 8, target; i >= 0; i-- {
		for c := cost[i]; (j-c == 0 || j-c > 0 && f[j-c] > 0) && f[j] == f[j-c]+1; j -= c {
			out = append(out, byte(i)+'1')
		}
	}
	return string(out)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
